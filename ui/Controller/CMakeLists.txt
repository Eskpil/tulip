set(Controller_SOURCES main.c arena.c packet.c)
set(Controller_HEADERS types.h packet.h arena.h)

pkg_check_modules(Systemd REQUIRED libsystemd)

add_executable(controller ${Controller_SOURCES} ${Controller_HEADERS} ${Systemd_HEADERS})

target_link_libraries(controller ${Systemd_LINK_LIBRARIES})
target_include_directories(controller PUBLIC ${Systemd_INCLUDE_DIRS})

target_compile_options(controller PUBLIC ${Systemd_CFLAGS_OTHER})
