#include <wayland-server.h>

#include "glass.h"
#include "input.h"
#include "output.h"
#include "view.h"

static void focus_view(struct glass_view *view,  struct wlr_surface *surface) {
  /* Note: this function only deals with keyboard focus. */
  if (view == NULL) {
    return;
  }
  struct glass_server *server = view->server;
  assert(server);

  struct wlr_seat *seat = server->seat;
  struct wlr_surface *prev_surface = seat->keyboard_state.focused_surface;
  if (prev_surface == surface) {
    /* Don't re-focus an already focused surface. */
    return;
  }
  if (prev_surface) {
    /*
		 * Deactivate the previously focused surface. This lets the client know
		 * it no longer has focus and the client will repaint accordingly, e.g.
		 * stop displaying a caret.
     */
    struct wlr_xdg_surface *previous = wlr_xdg_surface_try_from_wlr_surface(
        seat->keyboard_state.focused_surface);
    assert(previous != NULL && previous->role == WLR_XDG_SURFACE_ROLE_TOPLEVEL);
    wlr_xdg_toplevel_set_activated(previous->toplevel, false);
  }
  struct wlr_keyboard *keyboard = wlr_seat_get_keyboard(seat);
  /* Move the view to the front */

  if (view->kind == GLASS_VIEW_KIND_LAYER) {
    wlr_scene_node_raise_to_top(&view->layer.scene_tree->tree->node);
  } else {
    wlr_scene_node_raise_to_top(&view->xdg.scene_tree->node);
  }

  struct glass_output *output = glass_get_current_output(server);

  wl_list_remove(&view->link);
  wl_list_insert(&output->views, &view->link);
  /* Activate the new surface */
  if (view->kind == GLASS_VIEW_KIND_XDG) {
    wlr_xdg_toplevel_set_activated(view->xdg.xdg_toplevel, true);
  } else {
    // FIXME: What the heck to do with layer?!??!?!?!
  }
  /*
	 * Tell the seat to have the keyboard enter this surface. wlroots will keep
	 * track of this and automatically send key events to the appropriate
	 * clients without additional work on your part.
   */
  if (keyboard != NULL) {
    if (view->kind == GLASS_VIEW_KIND_LAYER) {
      wlr_seat_keyboard_notify_enter(
          seat,
          view->layer.surface->surface,
         keyboard->keycodes,
          keyboard->num_keycodes,
         &keyboard->modifiers
          );
  } else {
      wlr_seat_keyboard_notify_enter(
        seat,
        view->xdg.xdg_toplevel->base->surface,
        keyboard->keycodes,
        keyboard->num_keycodes,
        &keyboard->modifiers
        );
    }
  }
}

static void process_cursor_motion(struct glass_server *server, uint32_t time) {
  double sx, sy;
  struct wlr_seat *seat = server->seat;
  struct wlr_surface *surface = NULL;
  struct glass_view *view = glass_view_at(server,
                                             server->cursor->x, server->cursor->y, &surface, &sx, &sy);

  if (!view) {
    wlr_xcursor_manager_set_cursor_image(
        server->cursor_mgr, "default", server->cursor);
  }

  if (surface) {
    /*
		 * Send pointer enter and motion events.
		 *
		 * The enter event gives the surface "pointer focus", which is distinct
		 * from keyboard focus. You get pointer focus by moving the pointer over
		 * a window.
		 *
		 * Note that wlroots will avoid sending duplicate enter/motion events if
		 * the surface has already has pointer focus or if the client is already
		 * aware of the coordinates passed.
     */
    wlr_seat_pointer_notify_enter(seat, surface, sx, sy);
    wlr_seat_pointer_notify_motion(seat, time, sx, sy);
  } else {
    /* Clear pointer focus so future button events and such are not sent to
		 * the last client to have the cursor over it. */
    wlr_seat_pointer_clear_focus(seat);
  }

}

void glass_cursor_motion(struct wl_listener *listener, void *data) {
  struct glass_server *server =
      wl_container_of(listener, server, cursor_motion);
  struct wlr_pointer_motion_event *event = data;
  /* The cursor doesn't move unless we tell it to. The cursor automatically
	 * handles constraining the motion to the output layout, as well as any
	 * special configuration applied for the specific input device which
	 * generated the event. You can pass NULL for the device if you want to move
	 * the cursor around without any input. */
  wlr_cursor_move(server->cursor, &event->pointer->base,
                  event->delta_x, event->delta_y);

  process_cursor_motion(server, event->time_msec);
}

void glass_cursor_motion_absolute(struct wl_listener *listener, void *data) {
  struct glass_server *server =
      wl_container_of(listener, server, cursor_motion_absolute);
  struct wlr_pointer_motion_absolute_event *event = data;
  wlr_cursor_warp_absolute(server->cursor, &event->pointer->base, event->x,
                           event->y);

  process_cursor_motion(server, event->time_msec);
}

void glass_cursor_button(struct wl_listener *listener, void *data) {
  struct glass_server *server =
      wl_container_of(listener, server, cursor_button);
  struct wlr_pointer_button_event *event = data;
  /* Notify the client with pointer focus that a button press has occurred */
  wlr_seat_pointer_notify_button(server->seat,
                                 event->time_msec, event->button, event->state);

  double sx, sy;
  struct wlr_surface *surface = NULL;
  struct glass_view *view = glass_view_at(server,
                                             server->cursor->x, server->cursor->y, &surface, &sx, &sy);
  focus_view(view, surface);
}

void glass_cursor_axis(struct wl_listener *listener, void *data) {
  struct glass_server *server =
      wl_container_of(listener, server, cursor_axis);
  struct wlr_pointer_axis_event *event = data;
  /* Notify the client with pointer focus of the axis event. */
  wlr_seat_pointer_notify_axis(server->seat,
                               event->time_msec, event->orientation, event->delta,
                               event->delta_discrete, event->source);
}

void glass_cursor_frame(struct wl_listener *listener, void *data) {
  struct glass_server *server =
      wl_container_of(listener, server, cursor_frame);
  /* Notify the client with pointer focus of the frame event. */
  wlr_seat_pointer_notify_frame(server->seat);
}


void glass_seat_request_cursor(struct wl_listener *listener, void *data) {
  struct glass_server *server = wl_container_of(
      listener, server, request_cursor);

  struct wlr_seat_pointer_request_set_cursor_event *event = data;
  struct wlr_seat_client *focused_client =
      server->seat->pointer_state.focused_client;

  if (focused_client == event->seat_client) {
    wlr_cursor_set_surface(server->cursor, event->surface,
                           event->hotspot_x, event->hotspot_y);
  }
}

void glass_seat_request_set_selection(struct wl_listener *listener, void *data) {
  struct glass_server *server = wl_container_of(
      listener, server, request_set_selection);
  struct wlr_seat_request_set_selection_event *event = data;
  wlr_seat_set_selection(server->seat, event->source, event->serial);
}

static void keyboard_handle_modifiers(
    struct wl_listener *listener, void *data) {
  struct glass_keyboard *keyboard =
      wl_container_of(listener, keyboard, modifiers);

  wlr_seat_set_keyboard(keyboard->server->seat, keyboard->wlr_keyboard);
  wlr_seat_keyboard_notify_modifiers(keyboard->server->seat,
                                     &keyboard->wlr_keyboard->modifiers);
}


static void keyboard_handle_key(
    struct wl_listener *listener, void *data) {
  /* This event is raised when a key is pressed or released. */
  struct glass_keyboard *keyboard =
      wl_container_of(listener, keyboard, key);
  struct glass_server *server = keyboard->server;
  struct wlr_keyboard_key_event *event = data;
  struct wlr_seat *seat = server->seat;

  /* Otherwise, we pass it along to the client. */
  wlr_seat_set_keyboard(seat, keyboard->wlr_keyboard);
  wlr_seat_keyboard_notify_key(seat, event->time_msec,
                               event->keycode, event->state);
}


static void keyboard_handle_destroy(struct wl_listener *listener, void *data) {
  /* This event is raised by the keyboard base wlr_input_device to signal
	 * the destruction of the wlr_keyboard. It will no longer receive events
	 * and should be destroyed.
   */
  struct glass_keyboard *keyboard =
      wl_container_of(listener, keyboard, destroy);
  wl_list_remove(&keyboard->modifiers.link);
  wl_list_remove(&keyboard->key.link);
  wl_list_remove(&keyboard->destroy.link);
  wl_list_remove(&keyboard->link);
  free(keyboard);
}


static void glass_new_keyboard(struct glass_server *server,
                                struct wlr_input_device *device) {
  struct wlr_keyboard *wlr_keyboard = wlr_keyboard_from_input_device(device);

  struct glass_keyboard *keyboard =
      calloc(1, sizeof(struct glass_keyboard));

  keyboard->server = server;
  keyboard->wlr_keyboard = wlr_keyboard;

  /* We need to prepare an XKB keymap and assign it to the keyboard. This
	 * assumes the defaults (e.g. layout = "us"). */
  struct xkb_context *context = xkb_context_new(XKB_CONTEXT_NO_FLAGS);
  struct xkb_keymap *keymap = xkb_keymap_new_from_names(context, NULL,
                                                        XKB_KEYMAP_COMPILE_NO_FLAGS);

  wlr_keyboard_set_keymap(wlr_keyboard, keymap);
  xkb_keymap_unref(keymap);
  xkb_context_unref(context);
  wlr_keyboard_set_repeat_info(wlr_keyboard, 25, 600);

  /* Here we set up listeners for keyboard events. */
  keyboard->modifiers.notify = keyboard_handle_modifiers;
  wl_signal_add(&wlr_keyboard->events.modifiers, &keyboard->modifiers);
  keyboard->key.notify = keyboard_handle_key;
  wl_signal_add(&wlr_keyboard->events.key, &keyboard->key);
  keyboard->destroy.notify = keyboard_handle_destroy;
  wl_signal_add(&device->events.destroy, &keyboard->destroy);

  wlr_seat_set_keyboard(server->seat, keyboard->wlr_keyboard);

  /* And add the keyboard to our list of keyboards */
  wl_list_insert(&server->keyboards, &keyboard->link);
}



void glass_new_input_device(struct wl_listener *listener, void *data) {
  struct glass_server *server =
      wl_container_of(listener, server, new_input);
  struct wlr_input_device *device = data;
  switch (device->type) {
  case WLR_INPUT_DEVICE_POINTER:
    wlr_cursor_attach_input_device(server->cursor, device);
    break;
  case WLR_INPUT_DEVICE_KEYBOARD:
    glass_new_keyboard(server, device);
  default:
    break;
  }
  /* We need to let the wlr_seat know what our capabilities are, which is
	 * communiciated to the client. In TinyWL we always have a cursor, even if
	 * there are no pointer devices, so we always include that capability. */
  uint32_t caps = WL_SEAT_CAPABILITY_POINTER;
  if (!wl_list_empty(&server->keyboards)) {
    caps |= WL_SEAT_CAPABILITY_KEYBOARD;
  }

  wlr_seat_set_capabilities(server->seat, caps);
}
