#include <stdbool.h>
#include <stdlib.h>
#include <unistd.h>

#include <wayland-server-core.h>
#include <wlr/backend.h>
#include <wlr/render/allocator.h>
#include <wlr/render/wlr_renderer.h>
#include <wlr/types/wlr_cursor.h>
#include <wlr/types/wlr_compositor.h>
#include <wlr/types/wlr_data_device.h>
#include <wlr/types/wlr_output_layout.h>
#include <wlr/types/wlr_scene.h>
#include <wlr/types/wlr_seat.h>
#include <wlr/types/wlr_subcompositor.h>
#include <wlr/types/wlr_xcursor_manager.h>
#include <wlr/types/wlr_layer_shell_v1.h>
#include <wlr/types/wlr_xdg_shell.h>

#include <wlr/util/log.h>

#include "glass.h"
#include "input.h"
#include "output.h"
#include "view.h"

int main(void)
{
  wlr_log_init(WLR_DEBUG, NULL);

  struct glass_server server;

  setenv("WLR_X11_BACKEND_WINDOW_WIDTH", "800", false);
  setenv("WLR_X11_BACKEND_WINDOW_HEIGHT", "480", false);

  server.display = wl_display_create();
  server.backend = wlr_backend_autocreate(server.display, NULL);
  if (server.backend == NULL) {
    wlr_log(WLR_ERROR, "failed to create wlr_backend");
    return 1;
  }

  server.renderer = wlr_renderer_autocreate(server.backend);
  if (server.renderer == NULL) {
    wlr_log(WLR_ERROR, "failed to create wlr_renderer");
    return 1;
  }

  wlr_renderer_init_wl_display(server.renderer, server.display);
  server.allocator = wlr_allocator_autocreate(server.backend,
                                              server.renderer);
  if (server.allocator == NULL) {
    wlr_log(WLR_ERROR, "failed to create wlr_allocator");
    return 1;
  }

  const char *socket = wl_display_add_socket_auto(server.display);
  if (!socket) {
    wlr_backend_destroy(server.backend);
    return 1;
  }

  wlr_compositor_create(server.display, server.renderer);
  wlr_subcompositor_create(server.display);
  wlr_data_device_manager_create(server.display);

  server.output_layout = wlr_output_layout_create();

  wl_list_init(&server.outputs);
  server.new_output.notify = glass_new_output;
  wl_signal_add(&server.backend->events.new_output, &server.new_output);

  server.scene = wlr_scene_create();
  wlr_scene_attach_output_layout(server.scene, server.output_layout);

  server.seat = wlr_seat_create(server.display, "seat0");
  assert(server.seat);

  server.cursor = wlr_cursor_create();
  wlr_cursor_attach_output_layout(server.cursor, server.output_layout);

  server.cursor_mgr = wlr_xcursor_manager_create(NULL, 24);
  wlr_xcursor_manager_load(server.cursor_mgr, 1);
  server.cursor_mode = GLASS_CURSOR_PASSTHROUGH;
  server.cursor_motion.notify = glass_cursor_motion;
  wl_signal_add(&server.cursor->events.motion, &server.cursor_motion);
  server.cursor_motion_absolute.notify = glass_cursor_motion_absolute;
  wl_signal_add(&server.cursor->events.motion_absolute,
			&server.cursor_motion_absolute);
  server.cursor_button.notify = glass_cursor_button;
  wl_signal_add(&server.cursor->events.button, &server.cursor_button);
  server.cursor_axis.notify = glass_cursor_axis;
  wl_signal_add(&server.cursor->events.axis, &server.cursor_axis);
  server.cursor_frame.notify = glass_cursor_frame;
  wl_signal_add(&server.cursor->events.frame, &server.cursor_frame);


  server.new_input.notify = glass_new_input_device;
  wl_signal_add(&server.backend->events.new_input, &server.new_input);

  wl_list_init(&server.keyboards);
  server.request_cursor.notify = glass_seat_request_cursor;
  wl_signal_add(&server.seat->events.request_set_cursor,
                &server.request_cursor);
  server.request_set_selection.notify = glass_seat_request_set_selection;
  wl_signal_add(&server.seat->events.request_set_selection,
                &server.request_set_selection);


  server.layer_shell = wlr_layer_shell_v1_create(server.display, 3);
  wl_signal_add(&server.layer_shell->events.new_surface, &server.new_layered_surface);
  server.new_layered_surface.notify = glass_new_layered_surface;

  server.xdg_shell = wlr_xdg_shell_create(server.display, 3);
  server.new_xdg_surface.notify = glass_new_xdg_surface;
  wl_signal_add(&server.xdg_shell->events.new_surface,
                &server.new_xdg_surface);

  server.decoration_manager = wlr_xdg_decoration_manager_v1_create(server.display);
  server.new_toplevel_decoration.notify = glass_new_toplevel_decoration;
  wl_signal_add(&server.decoration_manager->events.new_toplevel_decoration, &server.new_toplevel_decoration);

  if (!wlr_backend_start(server.backend)) {
    wlr_backend_destroy(server.backend);
    wl_display_destroy(server.display);
    return 1;
  }

  /* Set the WAYLAND_DISPLAY environment variable to our socket and run the
	 * startup command if requested. */
  setenv("WAYLAND_DISPLAY", socket, true);
  /* Run the Wayland event loop. This does not return until you exit the
	 * compositor. Starting the backend rigged up all of the necessary event
	 * loop configuration to listen to libinput events, DRM events, generate
	 * frame events at the refresh rate, and so on. */

  if (fork() == 0) {
    execl("/bin/sh", "/bin/sh", "-c", "swaybg --image margot.jpg", (void *)NULL);
  }

  wlr_log(WLR_INFO, "sizeof(struct glass_view) = %lu", sizeof(struct glass_view));

  wlr_log(WLR_INFO, "Running Wayland compositor on WAYLAND_DISPLAY=%s",
          socket);
  wl_display_run(server.display);

  /* Once wl_display_run returns, we shut down the server. */
  wl_display_destroy_clients(server.display);
  wl_display_destroy(server.display);

  return 0;
}
