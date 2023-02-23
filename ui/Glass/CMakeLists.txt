set(Glass_Sources
        "${CMAKE_CURRENT_SOURCE_DIR}/main.c"
        "${CMAKE_CURRENT_SOURCE_DIR}/output.c"
        "${CMAKE_CURRENT_SOURCE_DIR}/view.c"
        "${CMAKE_CURRENT_SOURCE_DIR}/input.c")

pkg_check_modules(WaylandServer REQUIRED wayland-server)
pkg_check_modules(Wlroots REQUIRED wlroots)
pkg_check_modules(Pixman REQUIRED pixman-1)
pkg_check_modules(XkbCommon REQUIRED xkbcommon)

add_executable(glass ${Glass_Sources} ${WaylandServer_HEADERS} ${Wlroots_HEADERS} ${Pixman_HEADERS} ${XkbCommon_HEADERS})

target_link_libraries(glass LibProtocols ${WaylandServer_LINK_LIBRARIES} ${Wlroots_LINK_LIBRARIES} ${Pixman_LINK_LIBRARIES} ${XkbCommon_LINK_LIBRARIES})
target_include_directories(glass PUBLIC ${WaylandServer_INCLUDE_DIRS} ${Wlroots_INCLUDE_DIRS} ${Pixman_INCLUDE_DIRS} ${XkbCommon_INCLUDE_DIRS} ../Protocols)

target_compile_options(glass PUBLIC "-DWLR_USE_UNSTABLE" ${WaylandServer_CFLAGS_OTHER} ${Wlroots_CFLAGS_OTHER} ${Pixman_CFLAGS_OTHER} ${XkbCommon_CFLAGS_OTHER})
