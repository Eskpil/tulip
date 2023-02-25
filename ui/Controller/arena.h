#ifndef ARENA_H_
#define ARENA_H_

#include <stdlib.h>

#define ARENA_DEFAULT_CAPACITY (640 * 1000)
#define ARENA_DEFAULT_REGION_AMOUNT (1024)

struct region {
  struct region *next;
  size_t capacity;
  size_t size;
  char buffer[];
};

struct arena {
  struct region *first;
  struct region *last;
};

struct arena *arena_create();
void *arena_alloc(struct arena *arena, size_t size);
void *arena_realloc(struct arena *arena, void *old_ptr, size_t old_size, size_t new_size);
void arena_clean(struct arena *arena);
void arena_free(struct arena *arena);
void arena_summary(struct arena *arena);

#endif // ARENA_H_