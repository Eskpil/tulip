cmake_minimum_required(VERSION 3.22)
project(Tulip)

set(CMAKE_CXX_STANDARD 23)

add_compile_options(-Wno-address-of-packed-member)
add_compile_options(-Wcast-qual)
add_compile_options(-Wduplicated-cond)
add_compile_options(-Wformat=2)
add_compile_options(-Wimplicit-fallthrough)
add_compile_options(-Wlogical-op)
add_compile_options(-Wmisleading-indentation)
add_compile_options(-Wmissing-declarations)
add_compile_options(-Wno-unknown-warning-option)
add_compile_options(-Wundef)
add_compile_options(-Wunused)
add_compile_options(-Wno-unused-command-line-argument)
add_compile_options(-Wwrite-strings)
add_compile_options(-Wno-user-defined-literals)

add_compile_options("-fsanitize=address,undefined")
add_link_options("-fsanitize=address,undefined")

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wdeprecated-copy -Wnon-virtual-dtor -Wsuggest-override -Wno-literal-suffix")

find_package(PkgConfig REQUIRED)

include_directories(${CMAKE_SOURCE_DIR}/)

add_subdirectory(Protocols)
add_subdirectory(Utils)
add_subdirectory(Glass)