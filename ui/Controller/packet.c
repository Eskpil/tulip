#include <stdlib.h>
#include <assert.h>
#include <stdio.h>
#include <stdint.h>
#include <string.h>

#include <arpa/inet.h>

#include "packet.h"
#include "arena.h"

void dump_packet(struct controller_packet packet) {
  printf("\n");

  printf("  id: (%d)\n", packet.id);
  printf("  op: (%d)\n", packet.op);

  if (packet.op == OP_REQUEST) {
    printf("    request.version: (%d)\n", packet.request->version);
    printf("    request.hostname: (%s) (%ld)\n", packet.request->hostname.bytes, packet.request->hostname.size);
    printf("    request.kernel: (%s) (%ld)\n", packet.request->uname.bytes, packet.request->uname.size);
  } else {
    printf("    response.address.ipv4: (%s)\n", packet.response->address.bytes);
    printf("    response.private_key: (%s)\n", packet.response->private_key.bytes);
    printf("    response.public_key: (%s)\n", packet.response->public_key.bytes);
    printf("    response.port: (%d)\n", packet.response->port);
    printf("    response.version: (%d)\n", packet.response->version);
  }

  printf("\n");
}

static void encode_uint16(struct encoded_packet *packet, uint16_t value) {
  uint16_t sorted = htole16(value);
  packet->bytes[packet->cursor++] = (sorted >> 0) & 0xFF;
  packet->bytes[packet->cursor++] = (sorted >> 8) & 0xFF;
}

static void encode_uint32(struct encoded_packet *packet, uint32_t value) {
  uint32_t sorted = htole32(value);

  packet->bytes[packet->cursor++] = (sorted >> 0) & 0xFF;
  packet->bytes[packet->cursor++] = (sorted >> 8) & 0xFF;
  packet->bytes[packet->cursor++] = (sorted >> 16) & 0xFF;
  packet->bytes[packet->cursor++] = (sorted >> 24) & 0xFF;
}

static void encode_string(struct encoded_packet *packet, struct controller_string string) {
  assert(string.bytes);

  encode_uint32(packet, string.size);
  for (size_t i = 0; string.size > i; ++i) {
    packet->bytes[packet->cursor++] = string.bytes[i];
  }
}

static uint16_t decode_uint16(struct encoded_packet *packet) {
  uint8_t a = packet->bytes[packet->cursor++];
  uint8_t b = packet->bytes[packet->cursor++];

  uint16_t result = a + (b << 8);
  return result;
}

static uint32_t decode_uint32(struct encoded_packet *packet) {
  uint8_t a = packet->bytes[packet->cursor++];
  uint8_t b = packet->bytes[packet->cursor++];
  uint8_t c = packet->bytes[packet->cursor++];
  uint8_t d = packet->bytes[packet->cursor++];

  uint32_t result = a + (b << 8) + (c << 16) + (d << 24);
  return result;
}

static struct controller_string *decode_string(struct arena *arena, struct encoded_packet *packet) {
  struct controller_string *string = arena_alloc(arena, sizeof(struct controller_string));

  string->size = decode_uint32(packet);
  string->bytes = arena_alloc(arena, string->size + 10);

  for (size_t i = 0; string->size > i; ++i) {
    string->bytes[i] = packet->bytes[packet->cursor++];
  }

  return string;
}

struct encoded_packet *encode_packet(struct arena *arena, struct controller_packet packet) {
    struct encoded_packet *encoded = arena_alloc(arena, sizeof(struct encoded_packet));

    // FIXME: Find a better way of computing packet size. (considering strings)
    // this covers id and op.
    size_t size = sizeof(uint16_t) * 2;
    if (packet.op == OP_REQUEST) {
      size += sizeof(struct controller_request);

      size += packet.request->uname.size;
      size += packet.request->hostname.size;
    } else {
      size += sizeof(struct controller_response);
    }

    encoded->size = size;
    encoded->bytes = arena_alloc(arena, sizeof(u_int8_t) * size);

    encode_uint16(encoded, packet.id);
    encode_uint16(encoded, packet.op);

    if (packet.op == OP_REQUEST) {
      assert(packet.request);

      encode_uint16(encoded, packet.request->version);
      encode_string(encoded, packet.request->hostname);
      encode_string(encoded, packet.request->uname);
    } else {
      assert(packet.response);
      // FIXME: Implement encoding of OP_REQUEST
    }

    return encoded;
}

struct controller_packet *decode_packet(struct arena *arena, struct encoded_packet *encoded) {
  struct controller_packet *packet = arena_alloc(arena, sizeof(struct controller_packet));

  encoded->cursor = 0;

  packet->id = decode_uint16(encoded);
  packet->op = decode_uint16(encoded);

  if (packet->op == OP_REQUEST) {
    struct controller_request *request = arena_alloc(arena, sizeof(struct controller_request));
    request->version = decode_uint16(encoded);

    packet->request = request;
  } else {
    struct controller_response *response = arena_alloc(arena, sizeof(struct controller_response));

    response->address = *decode_string(arena, encoded);

    response->private_key = *decode_string(arena, encoded);
    response->public_key = *decode_string(arena, encoded);

    response->port = decode_uint16(encoded);
    response->version = decode_uint16(encoded);

    packet->response = response;
  }

  return packet;
}

