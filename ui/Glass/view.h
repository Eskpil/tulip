#ifndef VIEW_H_
#define VIEW_H_

#include <wlr/types/wlr_scene.h>
#include <wlr/types/wlr_layer_shell_v1.h>

#include "output.h"

enum glass_view_kind {
  GLASS_VIEW_KIND_XDG = 0,
  GLASS_VIEW_KIND_LAYER = 1,
};

/*
 *  NOTE:
 *
 *  Our view implementation does not host resizing, moving, maximization of xdg
 *  surfaces. This is because the compositor is in full control, the clients are
 *  only supposed to render. We control the rest.
 *
 */
struct glass_view {
  struct wl_list link;

  struct glass_server *server;
  struct glass_output *output;

  enum glass_view_kind kind;

  int x,y;

  union {
    struct {
      struct wlr_layer_surface_v1 *surface;
      struct wlr_scene_layer_surface_v1 *scene_tree;

      struct wl_listener map;
      struct wl_listener unmap;
      struct wl_listener destroy;
      struct wl_listener new_popup;
      struct wl_listener commit;
    } layer;

    struct {
      struct wlr_xdg_toplevel *xdg_toplevel;
      struct wlr_scene_tree *scene_tree;
      struct wl_listener map;
      struct wl_listener unmap;
      struct wl_listener destroy;
    } xdg;
  };
};

void glass_new_layered_surface(struct wl_listener *listener, void *data);
void glass_new_xdg_surface(struct wl_listener *listener, void *data);
void glass_new_toplevel_decoration(struct wl_listener *listener, void *data);

void glass_arrange_layers(struct glass_output *output);

struct glass_view *glass_view_at(struct glass_server *server, double lx, double ly, struct wlr_surface **surface, double *sx, double *sy);

#endif
