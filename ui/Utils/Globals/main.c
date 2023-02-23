#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <wayland-client.h>

static void global_registry_handler(void* data, struct wl_registry* registry, uint32_t id,
                                    const char* interface, uint32_t version)
{
  (void) data;
  (void) registry;

  printf(" \x1b[36m%-46s\x1b[0m \x1b[33m%-8d\x1b[0m \x1b[31m%-5d\x1b[0m\n", interface, version, id);
}

static void
global_registry_remover(void* data, struct wl_registry* registry, uint32_t id)
{
  // Don't care.

  (void) data;
  (void) registry;
  (void) id;
}

static const struct wl_registry_listener registry_listener = {
    global_registry_handler,
    global_registry_remover
};

int main(void)
{

  struct wl_display* display = wl_display_connect(NULL);

  if (display == NULL) {
    fprintf(stderr, "Can't connect to display\n");
    return 1;
  }

  printf("\n\x1b[36m%-46s\x1b[0m \x1b[33m%-8s\x1b[0m \x1b[31m%-5s\x1b[0m\n\n", "Interface:", "Version:", "Id:");
  struct wl_registry* registry = wl_display_get_registry(display);

  wl_registry_add_listener(registry, &registry_listener, NULL);

  wl_display_dispatch(display);

  wl_display_roundtrip(display);

  printf("\n");

  wl_display_disconnect(display);

  return 0;
}
