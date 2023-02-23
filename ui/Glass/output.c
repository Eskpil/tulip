#include <stdlib.h>

#include <wayland-server-core.h>
#include <wayland-util.h>

#include <wlr/types/wlr_output.h>
#include <wlr/types/wlr_output_damage.h>
#include <wlr/types/wlr_scene.h>
#include <wlr/types/wlr_output_layout.h>

#include "output.h"
#include "glass.h"

static void output_frame(struct wl_listener *listener, void *data) {
  (void) data;

  struct glass_output *output = wl_container_of(listener, output, frame);
  struct wlr_scene *scene = output->server->scene;

  struct wlr_scene_output *scene_output = wlr_scene_get_scene_output(
      scene, output->wlr_output);

  wlr_scene_output_commit(scene_output);

  struct timespec now;
  clock_gettime(CLOCK_MONOTONIC, &now);
  wlr_scene_output_send_frame_done(scene_output, &now);
}

static void output_request_state(struct wl_listener *listener, void *data) {
  struct glass_output *output = wl_container_of(listener, output, request_state);
  const struct wlr_output_event_request_state *event = data;
  wlr_output_commit_state(output->wlr_output, event->state);
}

static void output_destroy(struct wl_listener *listener, void *data) {
  (void) data;

  struct glass_output *output = wl_container_of(listener, output, destroy);

  wl_list_remove(&output->frame.link);
  wl_list_remove(&output->request_state.link);
  wl_list_remove(&output->destroy.link);
  wl_list_remove(&output->link);
  free(output);
}

void glass_new_output(struct wl_listener *listener, void *data) {
  struct glass_server *server =
      wl_container_of(listener, server, new_output);
  struct wlr_output *wlr_output = data;

  /* Configures the output created by the backend to use our allocator
	 * and our renderer. Must be done once, before commiting the output */
  wlr_output_init_render(wlr_output, server->allocator, server->renderer);

  /* Some backends don't have modes. DRM+KMS does, and we need to set a mode
	 * before we can use the output. The mode is a tuple of (width, height,
	 * refresh rate), and each monitor supports only a specific set of modes. We
	 * just pick the monitor's preferred mode, a more sophisticated compositor
	 * would let the user configure it. */
  if (!wl_list_empty(&wlr_output->modes)) {
    struct wlr_output_mode *mode = wlr_output_preferred_mode(wlr_output);
    wlr_output_set_mode(wlr_output, mode);
    wlr_output_enable(wlr_output, true);
    if (!wlr_output_commit(wlr_output)) {
      return;
    }
  }

  /* Allocates and configures our state for this output */
  struct glass_output *output =
      calloc(1, sizeof(struct glass_output));

  output->wlr_output = wlr_output;
  output->server = server;

  output->damage = wlr_output_damage_create(wlr_output);

  wl_list_init(&output->views);

  /* Sets up a listener for the frame event. */
  output->frame.notify = output_frame;
  wl_signal_add(&wlr_output->events.frame, &output->frame);

  /* Sets up a listener for the state request event. */
  output->request_state.notify = output_request_state;
  wl_signal_add(&wlr_output->events.request_state, &output->request_state);

  /* Sets up a listener for the destroy event. */
  output->destroy.notify = output_destroy;
  wl_signal_add(&wlr_output->events.destroy, &output->destroy);

  wl_list_insert(&server->outputs, &output->link);

  /* Adds this to the output layout. The add_auto function arranges outputs
	 * from left-to-right in the order they appear. A more sophisticated
	 * compositor would let the user configure the arrangement of outputs in the
	 * layout.
	 *
	 * The output layout utility automatically adds a wl_output global to the
	 * display, which Wayland clients can see to find out information about the
	 * output (such as DPI, scale factor, manufacturer, etc).
   */
  wlr_output_layout_add_auto(server->output_layout, wlr_output);
}

struct glass_output *glass_output_of_wlr_output(struct glass_server *server, struct wlr_output *wlr_output) {
  struct glass_output *output = NULL;

  wl_list_for_each(output, &server->outputs, link) {
    if (output->wlr_output == wlr_output) {
      return output;
    }
  }

  return NULL;
}

struct glass_output *glass_get_current_output(struct glass_server *server) {
  struct glass_output *output = NULL;

  wl_list_for_each(output, &server->outputs, link) {
    return output;
  }

  return NULL;
}