#include <stdio.h>
#include <string.h>
#include <errno.h>
#include <time.h>

#include <sys/epoll.h>

#include "arena.h"
#include "event.h"

struct controller_event_loop *event_loop_create(struct arena *arena) {
  struct controller_event_loop *event_loop = arena_alloc(arena, sizeof(struct controller_event_loop));

  event_loop->arena = arena;

  event_loop->epollfd = epoll_create1(0);
  if (event_loop->epollfd == -1) {
    fprintf(stderr, "[error]: Could not create epoll instance: (%s)\n", strerror(errno));
    exit(1);
  }

  event_loop->handlers = arena_alloc(arena, sizeof(struct controller_event_handlers));

  return event_loop;
}

void event_loop_insert(struct controller_event_loop *event_loop, struct controller_event event, controller_action_handler action_handler, void *data) {
  struct controller_event_handler *handler = arena_alloc(event_loop->arena, sizeof(struct controller_event_handler));

  srand (time(NULL) + event_loop->handlers->amount);
  handler->id = rand();

  handler->active = true;
  handler->event = event;
  handler->action = action_handler;
  handler->data = data;


  event_loop->handlers->amount += 1;
  struct controller_event_handler **new = arena_alloc(event_loop->arena, sizeof(struct controller_event_handler) * event_loop->handlers->amount);

  // Copy the old
  for (size_t i = 0; event_loop->handlers->amount > i; ++i) {
    if (i == event_loop->handlers->amount - 1) {
      new[i] = handler;
      continue;
    }

    new[i] = event_loop->handlers->data[i];
  }

  // Apply the new array
  event_loop->handlers->data = new;

  struct epoll_event epoll_event = {
      .data.u32 = handler->id,
      .events = event.events,
  };

  if (epoll_ctl(event_loop->epollfd, EPOLL_CTL_ADD, event.fd, &epoll_event) == -1) {
    fprintf(stderr, "[error]: epoll_ctl(%d, EPOLL_CTL_ADD, %d)", event_loop->epollfd, event.fd);
    return;
  }
}

static struct controller_event_handler *event_loop_find_handler(struct controller_event_loop *event_loop, uint32_t id) {
  for (size_t i = 0; event_loop->handlers->amount > i; ++i) {
    struct controller_event_handler *handler = event_loop->handlers->data[i];
    if (handler->id == id) {
      return handler;
    }
  }

  return NULL;
}

int event_loop_run(struct controller_event_loop *event_loop) {
#ifdef DEBUG
    for (size_t i = 0; event_loop->handlers->amount > i; ++i) {
      struct controller_event_handler *handler = event_loop->handlers->data[i];
      fprintf(stderr, "[info]: handler(%d, %d)\n", handler->id, handler->event.fd);
    }
#endif

    while (true) {
      if (event_loop->should_break) {
        break;
      }

      struct epoll_event *events = malloc(sizeof(struct epoll_event) * event_loop->handlers->amount);
      int nfds = epoll_wait(event_loop->epollfd, events, event_loop->handlers->amount, 1000);
      if (nfds == -1) {
        fprintf(stderr, "[error]: epoll_wait failed: (%s)", strerror(errno));
        continue;
      }

      if (nfds == 0) {
        continue;
      }

      for (size_t i = 0; nfds > i; ++i) {
        struct epoll_event event = events[i];

        struct controller_event_handler *handler = event_loop_find_handler(event_loop, event.data.u32);
        if (handler) {
          handler->action(event, handler->data);
        }
      }

      free(events);
    }

    return 0;
}
