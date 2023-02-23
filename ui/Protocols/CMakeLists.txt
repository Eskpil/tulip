set(PROTOCOL_SOURCES
        ${CMAKE_CURRENT_SOURCE_DIR}/wlr-layer-shell-unstable-v1-protocol.c
        ${CMAKE_CURRENT_SOURCE_DIR}/presentation-time.c
        ${CMAKE_CURRENT_SOURCE_DIR}/xdg-shell.c
        )

set(PROTOCOL_HEADERS
        ${CMAKE_CURRENT_SOURCE_DIR}/presentation-time-protocol.h
        ${CMAKE_CURRENT_SOURCE_DIR}/presentation-time-client.h

        ${CMAKE_CURRENT_SOURCE_DIR}/xdg-shell-protocol.h
        ${CMAKE_CURRENT_SOURCE_DIR}/xdg-shell-client.h

        ${CMAKE_CURRENT_SOURCE_DIR}/wlr-layer-shell-unstable-v1-protocol.h
        ${CMAKE_CURRENT_SOURCE_DIR}/wlr-layer-shell-unstable-v1-client.h)

add_library(LibProtocols ${PROTOCOL_SOURCES} ${PROTOCOL_HEADERS})