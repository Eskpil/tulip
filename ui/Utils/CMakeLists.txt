pkg_check_modules(WaylandClient REQUIRED wayland-client)

add_executable(globals globals.cc ${WaylandClient_HEADERS})

target_link_libraries(globals ${WaylandClient_LIBRARIES})
target_include_directories(globals PUBLIC ${WaylandClient_INCLUDE_DIRS})
target_compile_options(globals PUBLIC ${WaylandClient_CFLAGS_OTHER})
