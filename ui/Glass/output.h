#ifndef OUTPUT_H_
#define OUTPUT_H_

#include <wlr/types/wlr_output.h>
#include <wlr/types/wlr_output_damage.h>

struct glass_output {
  struct wl_list link;

  struct glass_server *server;
  struct wlr_output *wlr_output;

  struct wlr_output_damage *damage;

  struct wl_listener frame;
  struct wl_listener request_state;
  struct wl_listener destroy;

  struct wl_list views;

  struct {
    uint32_t left;
    uint32_t right;
    uint32_t top;
    uint32_t bottom;
  } excluded_margin;
};

void glass_new_output(struct wl_listener *listener, void *data);

struct glass_output *glass_output_of_wlr_output(struct glass_server *server, struct wlr_output *wlr_output);
struct glass_output *glass_get_current_output(struct glass_server *server);

#endif
