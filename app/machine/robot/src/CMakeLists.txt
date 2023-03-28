# Threads
find_package(Threads REQUIRED)
if (Threads_FOUND)
    message("")
endif ()

# Protobuf
find_package(Protobuf REQUIRED)
if (Protobuf_FOUND)
    message(STATUS "Protobuf_VERSION = ${Protobuf_VERSION}")
    message(STATUS "Protobuf_INCLUDE_DIRS = ${Protobuf_INCLUDE_DIRS}")
    message(STATUS "Protobuf_LIBRARY = ${Protobuf_LIBRARY}")
    message("")
endif ()

# gRPC
set(gRPC_SSL_PROVIDER package)
find_package(gRPC CONFIG REQUIRED)
if (gRPC_FOUND)
    message(STATUS "gRPC_VERSION = ${gRPC_VERSION}")
    message("")
endif ()

file(GLOB_RECURSE HEADER_FILES ${PROJECT_SOURCE_DIR}/include/*.h)
file(GLOB_RECURSE SRC_FILES
        *.cc
        *.h
        )

add_executable(${PROJECT_NAME} ${SRC_FILES})
target_include_directories(${PROJECT_NAME} PUBLIC ${PROJECT_SOURCE_DIR})
target_include_directories(${PROJECT_NAME} PUBLIC ${PROJECT_SOURCE_DIR}/src)
target_include_directories(${PROJECT_NAME} PUBLIC ${PROJECT_SOURCE_DIR}/include)
target_include_directories(${PROJECT_NAME} PUBLIC third_party/apis-cpp/src)
target_link_libraries(${PROJECT_NAME} Threads::Threads apis)