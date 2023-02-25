set(Controller_SOURCES main.c arena.c packet.c)
set(Controller_HEADERS arena.h packet.h types.h)

add_executable(controller ${Controller_SOURCES} ${Controller_HEADERS})
