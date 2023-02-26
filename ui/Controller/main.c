#include <stdio.h>
#include <string.h>
#include <errno.h>
#include <stdlib.h>

#include "arena.h"
#include "server.h"
#include "event.h"

int main(void) {
  struct arena *arena = arena_create();
  struct controller_event_loop *event_loop = event_loop_create(arena);
  struct controller_server *server = server_create(arena);

  if (server_bind(server) == -1) {
    fprintf(stderr, "[error]: Could not bind: (%s)\n", strerror(errno));
    exit(1);
  }

  server_send_request(server);

  {
    struct controller_event event = server_socket_event(server);
    event_loop_insert(event_loop, event, &server_socket_action, server);
  }

  {
    struct controller_event event = server_timer_event(server);
    event_loop_insert(event_loop, event, &server_timer_action, server);
  }

  event_loop_run(event_loop);

  arena_summary(arena);
  arena_free(arena);

  return 0;
}
