#include <assert.h>
#include <stdbool.h>
#include <stdlib.h>

#include <wayland-server-core.h>

#include <wlr/util/log.h>

#include <wlr/types/wlr_layer_shell_v1.h>
#include <wlr/types/wlr_compositor.h>
#include <wlr/types/wlr_xdg_shell.h>
#include <wlr/types/wlr_xdg_decoration_v1.h>

#include "view.h"
#include "glass.h"
#include "output.h"

static void on_view_commit(struct wl_listener *listener, void *data) {
  (void) data;

  struct glass_view *view = wl_container_of(listener, view, layer.commit);
  assert(view != NULL);

  wlr_log(WLR_INFO, "View: (%s) requested to be committed", view->layer.surface->namespace);
}

static void on_view_map(struct wl_listener *listener, void *data) {
  (void) data;

  struct glass_view *view = wl_container_of(listener, view, layer.map);
  assert(view);

  wlr_log(WLR_INFO, "View: (%s) requested to be mapped", view->layer.surface->namespace);
}

static void on_view_unmap(struct wl_listener *listener, void *data) {
  (void) data;

  struct glass_view *view = wl_container_of(listener, view, layer.unmap);
  assert(view != NULL);

  wlr_log(WLR_INFO, "View: (%s) requested to be unmapped", view->layer.surface->namespace);
}

static void on_view_destroy(struct wl_listener *listener, void *data) {
  (void) data;

  struct glass_view *view = wl_container_of(listener, view, layer.destroy);
  assert(view != NULL);

  wlr_log(WLR_INFO, "View: (%s) requested to be destroyed", view->layer.surface->namespace);
}

static void on_view_new_popup(struct wl_listener *listener, void *data) {
  (void) data;

  struct glass_view *view = wl_container_of(listener, view, layer.new_popup);
  assert(view != NULL);

  wlr_log(WLR_INFO, "View: (%s) requested a new popup", view->layer.surface->namespace);
}

void glass_new_layered_surface(struct wl_listener *listener, void *data) {
  struct glass_server *server =
      wl_container_of(listener, server, new_layered_surface);
  assert(server != NULL);

  struct wlr_layer_surface_v1 *surface = data;

  struct glass_view *view = calloc(1, sizeof(struct glass_view));
  view->kind = GLASS_VIEW_KIND_LAYER;
  view->server = server;

  view->layer.surface = surface;
  view->layer.scene_tree = wlr_scene_layer_surface_v1_create(&server->scene->tree, surface);
  view->layer.scene_tree->tree->node.data = view;

  view->layer.commit.notify = &on_view_commit;
  wl_signal_add(&surface->surface->events.commit, &view->layer.commit);

  view->layer.map.notify = &on_view_map;
  wl_signal_add(&surface->events.map, &view->layer.commit);

  view->layer.unmap.notify = &on_view_unmap;
  wl_signal_add(&surface->events.unmap, &view->layer.unmap);

  view->layer.destroy.notify = &on_view_destroy;
  wl_signal_add(&surface->events.destroy, &view->layer.destroy);

  view->layer.new_popup.notify = &on_view_new_popup;
  wl_signal_add(&surface->events.new_popup, &view->layer.new_popup);

  wlr_log(WLR_INFO, "New layered surface with namespace: (%s)", surface->namespace);

  struct glass_output *output = glass_output_of_wlr_output(server, surface->output);
  view->output = output;

  wl_list_insert(&output->views, &view->link);

  glass_arrange_layers(output);
}

static void on_xdg_toplevel_map(struct wl_listener *listener, void *data) {
  /* Called when the surface is mapped, or ready to display on-screen. */
  struct glass_view *view = wl_container_of(listener, view, xdg.map);
  wl_list_insert(&view->output->views, &view->link);
}

static void on_xdg_toplevel_unmap(struct wl_listener *listener, void *data) {
  /* Called when the surface is unmapped, and should no longer be shown. */
  struct glass_view *view = wl_container_of(listener, view, xdg.unmap);
  wl_list_remove(&view->link);
}

static void on_xdg_toplevel_destroy(struct wl_listener *listener, void *data) {
  /* Called when the surface is destroyed and should never be shown again. */
  struct glass_view *view = wl_container_of(listener, view, xdg.destroy);

  wl_list_remove(&view->xdg.map.link);
  wl_list_remove(&view->xdg.unmap.link);
  wl_list_remove(&view->xdg.destroy.link);

  free(view);
}

void glass_new_xdg_surface(struct wl_listener *listener, void *data) {
  struct glass_server *server = wl_container_of(listener, server, new_xdg_surface);
  struct wlr_xdg_surface *surface = data;

  assert(surface->role == WLR_XDG_SURFACE_ROLE_TOPLEVEL);

  struct glass_view *view = calloc(1, sizeof(struct glass_view));
  view->kind = GLASS_VIEW_KIND_XDG;
  view->server = server;

  view->xdg.xdg_toplevel = surface->toplevel;
  view->xdg.scene_tree = wlr_scene_xdg_surface_create(&server->scene->tree, surface);

  struct glass_output *output = glass_get_current_output(server);
  view->output = output;
  view->xdg.scene_tree->node.data = view;


  view->xdg.map.notify = on_xdg_toplevel_map;
  wl_signal_add(&surface->events.map, &view->xdg.map);
  view->xdg.unmap.notify = on_xdg_toplevel_unmap;
  wl_signal_add(&surface->events.unmap, &view->xdg.unmap);
  view->xdg.destroy.notify = on_xdg_toplevel_destroy;
  wl_signal_add(&surface->events.destroy, &view->xdg.destroy);

  wlr_log(WLR_INFO, "New xdg surface with title (%s)", surface->toplevel->title);

  wlr_xdg_toplevel_set_size(view->xdg.xdg_toplevel, output->wlr_output->width, output->wlr_output->height);
}

void glass_new_toplevel_decoration(struct wl_listener *listener, void *data) {
  struct glass_server *server = wl_container_of(listener, server, new_toplevel_decoration);
  struct wlr_xdg_toplevel_decoration_v1 *decoration = data;

  wlr_xdg_toplevel_decoration_v1_set_mode(decoration, WLR_XDG_TOPLEVEL_DECORATION_V1_MODE_SERVER_SIDE);
}

void glass_arrange_layers(struct glass_output *output) {
  uint32_t *margin_left = &output->excluded_margin.left;
  uint32_t *margin_right = &output->excluded_margin.right;
  uint32_t *margin_top = &output->excluded_margin.top;
  uint32_t *margin_bottom = &output->excluded_margin.bottom;

  *margin_left = 0;
  *margin_right = 0;
  *margin_top = 0;
  *margin_bottom = 0;

  uint32_t output_width = output->wlr_output->width;
  uint32_t output_height = output->wlr_output->height;

  // TODO: This layout function is inefficient and relies on some guesses about how the
  // layer shell is supposed to work, it needs reassessment later
  struct glass_view *view;
  wl_list_for_each_reverse(view, &output->views, link) {
    struct wlr_layer_surface_v1_state state = view->layer.surface->current;
    bool anchor_left = state.anchor & ZWLR_LAYER_SURFACE_V1_ANCHOR_LEFT;
    bool anchor_right = state.anchor & ZWLR_LAYER_SURFACE_V1_ANCHOR_RIGHT;
    bool anchor_top = state.anchor & ZWLR_LAYER_SURFACE_V1_ANCHOR_TOP;
    bool anchor_bottom = state.anchor & ZWLR_LAYER_SURFACE_V1_ANCHOR_BOTTOM;

    bool anchor_horiz = anchor_left && anchor_right;
    bool anchor_vert = anchor_bottom && anchor_top;

    uint32_t desired_width = state.desired_width;
    if (desired_width == 0) {
      desired_width = output_width;
    }
    uint32_t desired_height = state.desired_height;
    if (desired_height == 0) {
      desired_height = output_height;
    }

    uint32_t anchor_sum = anchor_left + anchor_right + anchor_top + anchor_bottom;
    switch (anchor_sum) {
    case 0:
      // Not anchored to any edge => display in centre with suggested size
      wlr_layer_surface_v1_configure(view->layer.surface, desired_width, desired_height);
      view->x = output_width / 2 - desired_width / 2;
      view->y = output_height / 2 - desired_height / 2;
      break;
    case 1:
      wlr_log(WLR_ERROR, "One anchor");
      // Anchored to one edge => use suggested size
      wlr_layer_surface_v1_configure(view->layer.surface,
                                     desired_width, desired_height);
      if (anchor_left || anchor_right) {
        view->y = output_height / 2 - desired_height / 2;
        if (anchor_left) {
          view->x = 0;
        } else {
          view->x = output_width - desired_width;
        }
      } else {
        view->x = output_width / 2 - desired_width / 2;
        if (anchor_top) {
          view->y = 0;
        } else {
          view->y = output_height - desired_height;
        }
      }
      wlr_log(WLR_ERROR, "Set layer surface x %d, y %d, width %d, height %d",
              view->x, view->y, desired_width, desired_height);
      break;
    case 2:
      // Anchored to two edges => use suggested size

      if (anchor_horiz) {
        wlr_layer_surface_v1_configure(view->layer.surface,
                                       desired_width, desired_height);
        view->x = 0;
        view->y = output_height / 2 - desired_height / 2;
      } else if (anchor_vert) {
        wlr_layer_surface_v1_configure(view->layer.surface,
                                       desired_width, desired_height);
        view->x = output_width / 2 - desired_width / 2;
        view->y = 0;
      } else if (anchor_top && anchor_left) {
        wlr_layer_surface_v1_configure(view->layer.surface,
                                       desired_width, desired_height);
        view->x = 0;
        view->y = 0;
      } else if (anchor_top && anchor_right) {
        wlr_layer_surface_v1_configure(view->layer.surface,
                                       desired_width, desired_height);
        view->x = output_width - desired_width;
        view->y = 0;
      } else if (anchor_bottom && anchor_right) {
        wlr_layer_surface_v1_configure(view->layer.surface,
                                       desired_width, desired_height);
        view->x = output_width - desired_width;
        view->y = output_height - desired_height;
      } else if (anchor_bottom && anchor_left) {
        wlr_layer_surface_v1_configure(view->layer.surface,
                                       desired_width, desired_height);
        view->x = 0;
        view->y = output_height - desired_height;
      }
      break;
    case 3:
      // Anchored to three edges => use suggested size on free axis only
      if (anchor_horiz) {
        wlr_layer_surface_v1_configure(view->layer.surface,
                                       output_width, desired_height);
        view->x = 0;
        if (anchor_top) {
          view->y = *margin_top;
          if (state.exclusive_zone) {
            *margin_top += desired_height;
          }
        } else {
          view->y = output_height - desired_height - *margin_bottom;
          if (state.exclusive_zone) {
            *margin_bottom += desired_height;
          }
        }
      } else {
        wlr_layer_surface_v1_configure(view->layer.surface,
                                       desired_width, output_height);
        view->y = 0;
        if (anchor_left) {
          view->x = *margin_left;
          if (state.exclusive_zone) {
            *margin_left += desired_width;
          }
        } else {
          view->x = output_width - desired_width - *margin_right;
          if (state.exclusive_zone) {
            *margin_right += desired_width;
          }
        }
      }
      break;
    case 4:
      // Fill the output
      wlr_layer_surface_v1_configure(view->layer.surface,
                                     output_width, output_height);
      view->x = 0;
      view->y = 0;
      break;
    default:
      UNREACHABLE()
    }
  }
}

struct glass_view *glass_view_at(
    struct glass_server *server, double lx, double ly,
    struct wlr_surface **surface, double *sx, double *sy) {

  struct wlr_scene_node *node = wlr_scene_node_at(
      &server->scene->tree.node, lx, ly, sx, sy);
  if (node == NULL || node->type != WLR_SCENE_NODE_BUFFER) {
    return NULL;
  }

  struct wlr_scene_buffer *scene_buffer = wlr_scene_buffer_from_node(node);
  struct wlr_scene_surface *scene_surface =
      wlr_scene_surface_try_from_buffer(scene_buffer);
  if (!scene_surface) {
    return NULL;
  }

  *surface = scene_surface->surface;
  struct wlr_scene_tree *tree = node->parent;
  while (tree != NULL && tree->node.data == NULL) {
    tree = tree->node.parent;
  }
  return tree->node.data;

}