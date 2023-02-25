#include <stdlib.h>
#include <assert.h>
#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include <stdint.h>

#include "arena.h"

static struct region *region_create(size_t capacity)
{
  const size_t region_size = sizeof(struct region) + capacity;
  struct region *region = malloc(region_size);
  memset(region, 0, region_size);
  region->capacity = capacity;
  return region;
}

static void *arena_alloc_aligned(struct arena *arena, size_t size, size_t alignment)
{
  if (arena->last == NULL) {
    assert(arena->first == NULL);

    struct region *region = region_create(
        size > ARENA_DEFAULT_CAPACITY ? size : ARENA_DEFAULT_CAPACITY);

    arena->last = region;
    arena->first = region;
  }

  // deal with the zero case here specially to simplify the alignment-calculating code below.
  if (size == 0) {
    // anyway, we now know we have *a* region -- so it's valid to just return it.
    return arena->last->buffer + arena->last->size;
  }

  // alignment must be a power of two.
  assert((alignment & (alignment - 1)) == 0);

  struct region *cur = arena->last;
  while (true) {

    char *ptr = (char*) (((uintptr_t) (cur->buffer + cur->size + (alignment - 1))) & ~(alignment - 1));
    size_t real_size = (size_t) ((ptr + size) - (cur->buffer + cur->size));

    if (cur->size + real_size > cur->capacity) {
      if (cur->next) {
        cur = cur->next;
        continue;
      } else {
        // out of space, make a new one. even though we are making a new region, there
        // aren't really any guarantees on the alignment of memory that malloc() returns.
        // so, allocate enough extra bytes to fix the 'worst case' alignment.
        size_t worst_case = size + (alignment - 1);

        struct region *region = region_create(worst_case > ARENA_DEFAULT_CAPACITY
                                               ? worst_case
                                               : ARENA_DEFAULT_CAPACITY);

        arena->last->next = region;
        arena->last = region;
        cur = arena->last;

        // ok, now we know we have enough space. just go back to the top of the loop here,
        // so we don't duplicate the code. we now know that we will definitely succeed,
        // so there won't be any infinite looping here.
        continue;
      }
    } else {
      memset(ptr, 0, real_size);
      cur->size += real_size;
      return ptr;
    }
  }
}

struct arena *arena_create()
{
  struct arena *arena = malloc(sizeof(struct arena));

  arena->first = NULL;
  arena->last = NULL;

  return arena;
}

void *arena_alloc(struct arena *arena, size_t size)
{
  // by default, align to a pointer size. this should be sufficient on most platforms.
  return arena_alloc_aligned(arena, size, sizeof(void*));
}

void *arena_realloc(struct arena *arena, void *old_ptr, size_t old_size, size_t new_size)
{
  if (old_size < new_size) {
    void *new_ptr = arena_alloc(arena, new_size);
    memcpy(new_ptr, old_ptr, old_size);
    return new_ptr;
  } else {
    return old_ptr;
  }
}

void arena_clean(struct arena *arena)
{
  for (struct region *iter = arena->first;
       iter != NULL;
       iter = iter->next) {
    iter->size = 0;
  }

  arena->last = arena->first;
}

void arena_free(struct arena *arena)
{
  struct region *iter = arena->first;
  while (iter != NULL) {
    struct region *next = iter->next;
    free(iter);
    iter = next;
  }
  arena->first = NULL;
  arena->last = NULL;

  free(arena);
}

void arena_summary(struct arena *arena)
{
  if (arena->first == NULL) {
    printf("[empty]");
  }

  for (struct region *iter = arena->first;
       iter != NULL;
       iter = iter->next) {
    printf("[%zu/%zu] -> ", iter->size, iter->capacity);
  }
  printf("\n");
}