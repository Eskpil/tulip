#ifndef SERVER_H_
#define SERVER_H_

#include <stdint.h>
#include <sys/epoll.h>

#include "event.h"

enum controller_transaction_state {
  CONTROLLER_TRANSACTION_STATE_NOT_FOUND = 0,
  CONTROLLER_TRANSACTION_STATE_REQUESTED = 1,
  CONTROLLER_TRANSACTION_STATE_RESPONDED = 2,
  CONTROLLER_TRANSACTION_STATE_LOST = 3,
  CONTROLLER_TRANSACTION_STATE_FINISHED = 4,
};

struct controller_transaction {
  uint16_t id;
  uint16_t attempts;
  enum controller_transaction_state state;
};

struct controller_transactions {
  uint32_t amount;
  struct controller_transaction **data;
};

struct controller_server {
  int sockfd;

  int timerfd;
  uint32_t timer_iterations;

  struct arena *arena;

  struct controller_transactions *transactions;
};

struct controller_server *server_create(struct arena *arena);

int server_bind(struct controller_server *server);
int server_fd(struct controller_server *server);

struct controller_event server_socket_event(struct controller_server *server);
int server_socket_action(struct epoll_event event, void *data);

struct controller_event server_timer_event(struct controller_server *server);
int server_timer_action(struct epoll_event event, void *data);

void server_send_request_with_id(struct controller_server *server, uint16_t id, bool insert_transaction);
void server_send_request(struct controller_server *server);

#endif //  SERVER_H_