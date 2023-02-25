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

#include "arena.h"
#include "packet.h"

struct controller_server {
  int sockfd;
  
  struct arena *arena;
};

enum controller_transaction_state {
  CONTROLLER_TRANSACTION_STATE_REQUESTED = 0,
  CONTROLLER_TRANSACTION_STATE_RESPONDED = 1,
  CONTROLLER_TRANSACTION_STATE_FINISHED = 2,
};

struct controller_transaction {
  uint16_t id;

};

static struct controller_server *server_create(struct arena *arena) {
  struct controller_server *server = arena_alloc(arena, sizeof(struct controller_server));

  server->arena = arena;
  server->sockfd = socket(AF_INET, SOCK_DGRAM, 0);

  if (server->sockfd == -1) {
    fprintf(stderr, "[error]: Could not create socket: (%s)\n", strerror(errno));
    exit(1);
  }

  return server;
}

static int server_fd(struct controller_server *server) {
  return server->sockfd;
}

static int server_action(struct controller_server *server, struct epoll_event event) {
  if (event.events == EPOLLIN) {
    uint8_t *buffer = arena_alloc(server->arena, PACKET_MAX_SIZE);

    struct sockaddr_in client_addr;

    socklen_t len = sizeof(struct sockaddr_in);
    int nread = recvfrom(server->sockfd, buffer, PACKET_MAX_SIZE, MSG_WAITALL, (struct sockaddr *)&client_addr, &len);
    if (nread == -1) {
      fprintf(stderr, "[error]: Could not read from socket: (%s)\n", strerror(errno));
      return nread;
    }

    char addr_as_str[INET_ADDRSTRLEN];
    inet_ntop(AF_INET, &(client_addr.sin_addr.s_addr), addr_as_str, INET_ADDRSTRLEN);

    fprintf(stderr, "[info]: Got connection from: (%s:%d)\n", addr_as_str, client_addr.sin_port);

    struct encoded_packet *encoded_packet = arena_alloc(server->arena, sizeof(struct encoded_packet));
    encoded_packet->bytes = arena_alloc(server->arena, nread);

    memcpy(encoded_packet->bytes, buffer, nread);
    encoded_packet->size = nread;

    struct controller_packet *packet = decode_packet(server->arena, encoded_packet);

    printf("decoded: \n");
    dump_packet(*packet);
  }

  return 0;
}

static struct epoll_event server_event(struct controller_server *server) {
  struct epoll_event event = {
      .data.fd = server->sockfd,
      .events = EPOLLIN,
  };

  return event;
}

static int server_bind(struct controller_server *server) {
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

static void server_send_request(struct controller_server *server, struct arena *arena) {
  struct controller_packet *packet = arena_alloc(arena, sizeof(struct controller_packet));
  packet->request = arena_alloc(arena, sizeof(struct controller_request));

  srand (time(NULL));

  packet->id = rand();
  packet->request->version = 12;

  struct utsname *utsname = arena_alloc(arena, sizeof(struct utsname));
  if (uname(utsname) == -1) {
    fprintf(stderr, "[error]: could not uname: (%s)\n", strerror(errno));
    return;
  }

  char *kernel = arena_alloc(arena, 255);
  snprintf(kernel, 255, "%s %s", utsname->sysname, utsname->release);

  packet->request->uname.bytes = arena_alloc(arena, strlen(kernel));
  packet->request->uname.size = strlen(kernel);
  memcpy(packet->request->uname.bytes, kernel, 255);

  char hostname[63];
  if (gethostname(hostname, 63) == -1) {
    fprintf(stderr, "[error]: could not get hostname: (%s)\n", strerror(errno));
    return;
  }

  packet->request->hostname.bytes = arena_alloc(arena, strlen(hostname));
  packet->request->hostname.size = strlen(hostname);
  memcpy(packet->request->hostname.bytes, hostname, strlen(hostname));

  dump_packet(*packet);

  int sockfd = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);
  if (sockfd == -1) {
    fprintf(stderr, "[error]: Failed to create socket: (%s)\n", strerror(errno));
    return;
  }

  struct encoded_packet *encoded = encode_packet(arena, *packet);

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
}

int main(void) {
  struct arena *arena = arena_create();

  int epollfd = epoll_create1(0);
  if (epollfd == -1) {
    fprintf(stderr, "[error]: Could not create epoll instance: (%s)\n", strerror(errno));
    exit(1);
  }

  struct controller_server *server = server_create(arena);

  if (server_bind(server) == -1) {
    fprintf(stderr, "[error]: Could not bind: (%s)\n", strerror(errno));
    exit(1);
  }

  server_send_request(server, arena);

  {
    struct epoll_event event = server_event(server);
    if (epoll_ctl(epollfd, EPOLL_CTL_ADD, event.data.fd, &event) == -1) {
      fprintf(stderr, "[error]: failed to epoll_ctl in server: (%s)", strerror(errno));
      exit(1);
    }
  }

  while (true) {
    struct epoll_event *events = malloc(sizeof(struct epoll_event) * 10);
    int nfds = epoll_wait(epollfd, events, 10, 1000);
    if (nfds == -1) {
      fprintf(stderr, "[error]: epoll_wait failed: (%s)", strerror(errno));
      continue;
    }

    if (nfds == 0) {
      continue;
    }

    fprintf(stderr, "[info]: eventloop awake (%d)\n", nfds);

    for (size_t i = 0; nfds > i; ++i) {
      struct epoll_event event = events[i];

      if (event.data.fd == server_fd(server)) {
        server_action(server, event);
      }
    }

    free(events);
  }

  arena_summary(arena);
  arena_free(arena);

  return 0;
}
