#ifndef PACKET_H_
#define PACKET_H_

#include <stdint.h>
#include <stdbool.h>
#include <arpa/inet.h>

#include "arena.h"
#include "types.h"

#define PACKET_MAX_SIZE 1024 * 4

enum controller_request_op {
  OP_REQUEST = 0,
  OP_RESPONSE = 1,
};

struct controller_service {
  struct controller_string name;
  uint16_t port;
};

struct controller_services {
  uint32_t amount;
  struct controller_service *data;
};

struct controller_request {
  uint16_t version;
  struct controller_string device_class;
  struct controller_string hostname;
  struct controller_string uname;
};

struct controller_response {
  struct controller_string address;

  struct controller_string private_key;
  struct controller_string public_key;

  struct controller_services *services;

  uint16_t version;
};

struct controller_packet {
  uint16_t id;
  enum controller_request_op op;

  union {
    struct controller_request *request;
    struct controller_response *response;
  };
};

struct encoded_packet {
  uint8_t *bytes;
  size_t size;
  size_t cursor;
};

struct encoded_packet *encode_packet(struct arena *arena, struct controller_packet packet);
struct controller_packet *decode_packet(struct arena *arena, struct encoded_packet *packet);
void dump_packet(struct controller_packet packet);

#endif
