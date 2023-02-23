#ifndef GLASS_H_
#define GLASS_H_

#include <assert.h>
#include <getopt.h>
#include <stdbool.h>
#include <stdlib.h>
#include <stdio.h>
#include <time.h>
#include <unistd.h>
#include <wayland-server-core.h>
#include <wlr/backend.h>
#include <wlr/render/allocator.h>
#include <wlr/render/wlr_renderer.h>
#include <wlr/types/wlr_cursor.h>
#include <wlr/types/wlr_compositor.h>
#include <wlr/types/wlr_data_device.h>
#include <wlr/types/wlr_input_device.h>
#include <wlr/types/wlr_keyboard.h>
#include <wlr/types/wlr_output.h>
#include <wlr/types/wlr_output_layout.h>
#include <wlr/types/wlr_pointer.h>
#include <wlr/types/wlr_scene.h>
#include <wlr/types/wlr_seat.h>
#include <wlr/types/wlr_subcompositor.h>
#include <wlr/types/wlr_xcursor_manager.h>
#include <wlr/types/wlr_layer_shell_v1.h>
#include <wlr/types/wlr_xdg_shell.h>
#include <wlr/types/wlr_xdg_decoration_v1.h>

#include <wlr/util/log.h>
#include <xkbcommon/xkbcommon.h>

#define EXIT_WITH_MESSAGE(MESSAGE)                  \
    wlr_log(WLR_ERROR, "%s", MESSAGE);              \
    exit(EXIT_FAILURE);

#define EXIT_WITH_FORMATTED_MESSAGE(MESSAGE, ...)   \
    wlr_log(WLR_ERROR, MESSAGE, __VA_ARGS__);       \
    exit(EXIT_FAILURE);

#define UNREACHABLE() \
    EXIT_WITH_MESSAGE("UNREACHABLE");

enum glass_cursor_mode {
  GLASS_CURSOR_PASSTHROUGH,
  GLASS_CURSOR_MOVE,
  GLASS_CURSOR_RESIZE,
};

struct glass_server {
  struct wl_display *display;
  struct wlr_backend *backend;

  struct wlr_renderer *renderer;
  struct wlr_allocator *allocator;
  struct wlr_scene *scene;

  struct wlr_layer_shell_v1 *layer_shell;
  struct wl_listener new_layered_surface;

  struct wlr_xdg_decoration_manager_v1 *decoration_manager;
  struct wl_listener new_toplevel_decoration;

  struct wlr_xdg_shell *xdg_shell;
  struct wl_listener new_xdg_surface;

  struct wl_list views;

  enum glass_cursor_mode cursor_mode;
  struct wlr_cursor *cursor;
  struct wlr_xcursor_manager *cursor_mgr;

  struct wl_listener cursor_motion;
  struct wl_listener cursor_motion_absolute;
  struct wl_listener cursor_button;
  struct wl_listener cursor_axis;
  struct wl_listener cursor_frame;

  struct wlr_seat *seat;
  struct wl_listener new_input;
  struct wl_listener request_cursor;
  struct wl_listener request_set_selection;
  struct wl_list keyboards;

  struct wlr_output_layout *output_layout;
  struct wl_list outputs;
  struct wl_listener new_output;
};

#endif
