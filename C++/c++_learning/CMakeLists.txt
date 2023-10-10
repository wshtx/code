cmake_minimum_required(VERSION 3.22)
project(c++_learning)


set(CMAKE_CXX_STANDARD 17)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${PROJECT_SOURCE_DIR}/build)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${PROJECT_SOURCE_DIR}/build)

add_executable(test main.cpp)

target_link_directories(test PRIVATE .)

file(GLOB srcs *.cpp)
target_sources(test PRIVATE ${srcs})
