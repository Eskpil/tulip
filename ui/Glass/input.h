#ifndef CURSOR_H_
#define CURSOR_H_

struct glass_keyboard {
  struct wl_list link; // glass_server::keyboards

  struct glass_server *server;
  struct wlr_keyboard *wlr_keyboard;

  struct wl_listener modifiers;
  struct wl_listener key;
  struct wl_listener destroy;
};

void glass_cursor_motion(struct wl_listener *listener, void *data);
void glass_cursor_motion_absolute(struct wl_listener *listener, void *data);
void glass_cursor_button(struct wl_listener *listener, void *data);
void glass_cursor_axis(struct wl_listener *listener, void *data);
void glass_cursor_frame(struct wl_listener *listener, void *data);

void glass_seat_request_cursor(struct wl_listener *listener, void *data);
void glass_seat_request_set_selection(struct wl_listener *listener, void *data);

void glass_new_input_device(struct wl_listener *listener, void *data);

#endif