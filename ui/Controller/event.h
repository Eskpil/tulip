#ifndef EVENT_H_
#define EVENT_H_

#include <stdbool.h>

#include <sys/epoll.h>

#include "arena.h"

typedef int (*controller_action_handler)(struct epoll_event event, void *data);

struct controller_event {
  int fd;
  int events;
};

struct controller_event_handler {
  uint32_t id;
  bool active;
  struct controller_event event;
  void *data;
  controller_action_handler action;
};

struct controller_event_handlers {
  uint32_t amount;
  struct controller_event_handler **data;
};

struct controller_event_loop {
  struct arena *arena;

  bool should_break;

  int epollfd;
  struct controller_event_handlers *handlers;
};

struct controller_event_loop *event_loop_create(struct arena *arena);
void event_loop_insert(struct controller_event_loop *event_loop, struct controller_event event, controller_action_handler handler, void *data);

int event_loop_run(struct controller_event_loop *event_loop);

#endif // EVENT_H_