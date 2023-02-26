#include <assert.h>
#include <stdio.h>
#include <string.h>
#include <errno.h>
#include <stdlib.h>
#include <stddef.h>
#include <stdbool.h>
#include <time.h>
#include <unistd.h>

#include <sys/utsname.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#include <sys/epoll.h>
#include <sys/timerfd.h>

#include "arena.h"
#include "server.h"
#include "packet.h"

struct controller_server *server_create(struct arena *arena) {
  struct controller_server *server = arena_alloc(arena, sizeof(struct controller_server));

  server->arena = arena;
  server->transactions = arena_alloc(server->arena, sizeof(struct controller_transactions));

  server->sockfd = socket(AF_INET, SOCK_DGRAM, 0);

  if (server->sockfd == -1) {
    fprintf(stderr, "[error]: Could not create socket: (%s)\n", strerror(errno));
    exit(1);
  }

  server->timerfd = timerfd_create(CLOCK_REALTIME, 0);
  if (server->timerfd == -1) {
    fprintf(stderr, "[error]: could not create timerfd: (%s)\n", strerror(errno));
    exit(1);
  }

  // Create a repeating timer that executes every 10 seconds.
  // NOTE: The timer will be used to try recover unanswered transactions.
  struct itimerspec ispec = {
      .it_value.tv_sec = 10,
      .it_interval.tv_sec = 10,
  };

  if (timerfd_settime(server->timerfd, TFD_TIMER_ABSTIME, &ispec, NULL) == -1) {
    fprintf(stderr, "[error]: could not set time of timerfd: (%s)\n", strerror(errno));
    exit(1);
  }

  return server;
}

int server_fd(struct controller_server *server) {
  return server->sockfd;
}

static struct controller_transaction *server_find_transaction(struct controller_server *server, uint16_t id) {
  for (size_t i = 0; server->transactions->amount > i; ++i) {
    struct controller_transaction *transaction = server->transactions->data[i];
    if (transaction->id == id) {
      return transaction;
    }
  }

  return NULL;
}

static void server_insert_transaction(struct controller_server *server, struct controller_transaction *transaction) {
  server->transactions->amount += 1;
  struct controller_transaction **new = arena_alloc(server->arena, sizeof(struct controller_transaction) * server->transactions->amount);

  // Copy the old
  for (size_t i = 0; server->transactions->amount > i; ++i) {
    if (i == server->transactions->amount - 1) {
      new[i] = transaction;
      continue;
    }

    new[i] = server->transactions->data[i];
  }

  // Apply the new array
  server->transactions->data = new;
}

int server_socket_action(struct epoll_event event, void *data) {
  struct controller_server *server = data;

  if (event.events & EPOLLIN) {
    uint8_t *buffer = malloc(PACKET_MAX_SIZE);

    struct sockaddr_in client_addr;

    socklen_t len = sizeof(struct sockaddr_in);
    int nread = recvfrom(server->sockfd, buffer, PACKET_MAX_SIZE, MSG_WAITALL, (struct sockaddr *)&client_addr, &len);
    if (nread == -1) {
      fprintf(stderr, "[error]: Could not read from socket: (%s)\n", strerror(errno));
      free(buffer);
      return nread;
    }

    char addr_as_str[INET_ADDRSTRLEN];
    inet_ntop(AF_INET, &(client_addr.sin_addr.s_addr), addr_as_str, INET_ADDRSTRLEN);

    fprintf(stderr, "[info]: Got connection from: (%s:%d)\n", addr_as_str, client_addr.sin_port);

    struct encoded_packet *encoded_packet = arena_alloc(server->arena, sizeof(struct encoded_packet));
    encoded_packet->bytes = arena_alloc(server->arena, nread);

    memset(encoded_packet->bytes, '\0', PACKET_MAX_SIZE);

    memcpy(encoded_packet->bytes, buffer, nread);
    encoded_packet->size = nread;

    free(buffer);

    struct controller_packet *packet = decode_packet(server->arena, encoded_packet);
    assert(packet);

    struct controller_transaction *transaction = server_find_transaction(server, packet->id);
    if (!transaction) {
      fprintf(stderr, "[error]: No transaction with id: (%d)\n", packet->id);
      return 1;
    }

    fprintf(stderr, "[info]: Found transaction for id: (%d)\n", transaction->id);
    transaction->state = CONTROLLER_TRANSACTION_STATE_RESPONDED;
  }

  return 0;
}

struct controller_event server_socket_event(struct controller_server *server) {
  struct controller_event event = {
      .fd = server_fd(server),
      .events = EPOLLIN,
  };

  return event;
}

struct controller_event server_timer_event(struct controller_server *server) {
  struct controller_event event = {
      .fd = server->timerfd,
      .events = EPOLLIN | EPOLLET,
  };

  return event;
}

int server_timer_action(struct epoll_event event, void *data) {
  struct controller_server *server = data;

  if (event.events & EPOLLIN) {
    uint64_t value;
    read(server->timerfd, &value, 8);

    if (server->timer_iterations == 0) {
      server->timer_iterations += 1;
      return 0;
    }

    for (size_t i = 0; server->transactions->amount > i; ++i) {
      struct controller_transaction *transaction = server->transactions->data[i];
      if (transaction->state == CONTROLLER_TRANSACTION_STATE_REQUESTED) {

        if (transaction->attempts > 10) {
          transaction->state = CONTROLLER_TRANSACTION_STATE_LOST;
          fprintf(stderr, "[info]: giving up on transaction: (%d) after (%d) attempts\n", transaction->id, transaction->attempts);
          continue;
        }

        server_send_request_with_id(server, transaction->id, false);
        transaction->attempts += 1;
      }
    }
  }

  server->timer_iterations += 1;
}

int server_bind(struct controller_server *server) {
  struct sockaddr_in addr = {0};

  addr.sin_addr.s_addr = INADDR_ANY;
  addr.sin_family = AF_INET;
  addr.sin_port = htons(6543);

  int rc = bind(server->sockfd, (struct sockaddr *)&addr, sizeof(struct sockaddr_in));
  if (rc == -1) {
    return rc;
  }

  char addr_as_str[INET_ADDRSTRLEN];
  inet_ntop(AF_INET, &(addr.sin_addr.s_addr), addr_as_str, INET_ADDRSTRLEN);

  fprintf(stderr, "[info]: listening on: (%s:6543)\n", addr_as_str);

  return 0;
}

void server_send_request_with_id(struct controller_server *server, uint16_t id, bool insert_transaction) {
  struct controller_packet *packet = arena_alloc(server->arena, sizeof(struct controller_packet));

  packet->id = id;
  packet->request = arena_alloc(server->arena, sizeof(struct controller_request));

  packet->request->version = 12;

  struct utsname *utsname = arena_alloc(server->arena, sizeof(struct utsname));
  if (uname(utsname) == -1) {
    fprintf(stderr, "[error]: could not uname: (%s)\n", strerror(errno));
    return;
  }

  char *kernel = arena_alloc(server->arena, 255);
  snprintf(kernel, 255, "%s %s", utsname->sysname, utsname->release);

  packet->request->uname.bytes = arena_alloc(server->arena, strlen(kernel));
  packet->request->uname.size = strlen(kernel);
  memcpy(packet->request->uname.bytes, kernel, 255);

  char hostname[63];
  if (gethostname(hostname, 63) == -1) {
    fprintf(stderr, "[error]: could not get hostname: (%s)\n", strerror(errno));
    return;
  }

  packet->request->hostname.bytes = arena_alloc(server->arena, strlen(hostname));
  packet->request->hostname.size = strlen(hostname);
  memcpy(packet->request->hostname.bytes, hostname, strlen(hostname));

  packet->request->device_class.bytes = arena_alloc(server->arena, strlen("interface"));
  packet->request->device_class.size = strlen("interface");
  sprintf(packet->request->device_class.bytes, "interface");

  int sockfd = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);
  if (sockfd == -1) {
    fprintf(stderr, "[error]: Failed to create socket: (%s)\n", strerror(errno));
    return;
  }

  struct encoded_packet *encoded = encode_packet(server->arena, *packet);

#ifdef DEBUG
  fprintf(stderr, "[");
  for(size_t i = 0; encoded->size > i; ++i) {
    fprintf(stderr, "%d, ", encoded->bytes[i]);
  }
  fprintf(stderr, "]\n");
#endif

  int yes = 1;
  int rc = setsockopt(sockfd, SOL_SOCKET, SO_BROADCAST, (char*)&yes, sizeof(yes));
  if (rc == -1) {
    fprintf(stderr, "[error]: setsockopt(SO_BROADCAST): (%s)", strerror(errno));
    return;
  }


  struct sockaddr_in dst;
  memset(&dst, 0, sizeof(dst));
  dst.sin_family = AF_INET;
  dst.sin_port = htons(7654);
  dst.sin_addr.s_addr = INADDR_BROADCAST;

  rc = sendto(sockfd, encoded->bytes, encoded->size, 0, (struct sockaddr*)&dst, sizeof(dst));
  if (rc == -1) {
    fprintf(stderr, "[errno]: could not send: (%s)\n", strerror(errno));
    return;
  }

  close(sockfd);

  if (insert_transaction) {
    struct controller_transaction *transaction = arena_alloc(server->arena, sizeof(struct controller_transaction));

    transaction->id = packet->id;
    transaction->state = CONTROLLER_TRANSACTION_STATE_REQUESTED;

    server_insert_transaction(server, transaction);
  }
}

void server_send_request(struct controller_server *server) {
  srand (time(NULL));
  uint16_t id = rand();

  return server_send_request_with_id(server, id, true);
}