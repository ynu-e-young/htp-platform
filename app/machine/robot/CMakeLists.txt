cmake_minimum_required(VERSION 3.10)

project(htp-platform.machine.robot
        LANGUAGES C CXX
        HOMEPAGE_URL "https://github.com/hominsu/htp-platform")

# setting definitions
add_definitions(-DPROJECT_NAME=\"${PROJECT_NAME}\")
add_definitions(-DHOMEPAGE_URL=\"${PROJECT_HOMEPAGE_URL}\")
message("")
MESSAGE(STATUS "Home page url: ${PROJECT_HOMEPAGE_URL}")

set(CMAKE_CXX_STANDARD 17)


# ============================== select operating system  ==============================
message("")
message("Operation system is ${CMAKE_SYSTEM}")
message("Current compiler: ${CMAKE_CXX_COMPILER_ID}")
if (CMAKE_SYSTEM_NAME MATCHES "Darwin")
    link_directories(/opt/homebrew/lib)
    include_directories(/opt/homebrew/include)
    message(STATUS "Load link directories from /opt/homebrew/lib")
    message(STATUS "Load include directories from /opt/homebrew/include")
    set(OPENSSL_ROOT_DIR /opt/homebrew/Cellar/openssl@1.1/1.1.1q)
    set(OPENSSL_INCLUDE_DIR /opt/homebrew/Cellar/openssl@1.1/1.1.1q/include)
elseif (CMAKE_SYSTEM_NAME MATCHES "Linux")
    link_directories(/usr/local/lib)
    include_directories(/usr/local/include)
    message(STATUS "Load link directories from /usr/local/lib")
    message(STATUS "Load include directories from /usr/local/include")
else ()
    message(FATAL_ERROR "Platform ${CMAKE_SYSTEM} is not support for this project")
endif ()


# ============================== find_package ==============================
# Threads
find_package(Threads REQUIRED)
if (Threads_FOUND)
    message("")
endif ()

# OpenSSL
set(OPENSSL_USE_STATIC_LIBS true)
find_package(OpenSSL REQUIRED)
if (OPENSSL_FOUND)
    message(STATUS "OPENSSL_VERSION = ${OPENSSL_VERSION}")
    message(STATUS "OPENSSL_INCLUDE_DIR = ${OPENSSL_INCLUDE_DIR}")
    message(STATUS "OPENSSL_LIBRARIES = ${OPENSSL_LIBRARIES}")
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
find_package(gRPC CONFIG REQUIRED)
if (gRPC_FOUND)
    message(STATUS "gRPC_VERSION = ${gRPC_VERSION}")
    message("")
endif ()


# ============================== add executable and link libraries ==============================
include_directories(${PROJECT_SOURCE_DIR})
include_directories(${PROJECT_SOURCE_DIR}/include)
include_directories(${PROJECT_SOURCE_DIR}/../../../)
file(GLOB ProjectSRC
        "../../../api/machine/service/v1/cpp/*.cc"
        "../../../api/machine/robot/v1/cpp/*.cc"
        "include/robot/internal/data/motor/motor_config/*.cc"
        "include/robot/internal/data/motor/serial/*.cc"
        "include/robot/internal/data/motor/*.cc"
        "include/robot/internal/data/grpc/*.cc"
        "include/robot/internal/service/*.cc"
        "include/robot/internal/server/*.cc"
        "include/robot/internal/utils/*.cc"
        "include/robot/*.cc"
        "include/conf/*.cc"
        "*.cc")

if (CMAKE_BUILD_TYPE STREQUAL Debug)
    ADD_DEFINITIONS(-DDEBUG)
    message(STATUS "CMake Build Type: Debug")
    message("")
elseif (CMAKE_BUILD_TYPE STREQUAL Release)
    ADD_DEFINITIONS(-DRELEASE)
    message(STATUS "CMake Build Type: Release")
    message("")
endif ()

add_executable(${PROJECT_NAME} ${ProjectSRC})
target_link_libraries(${PROJECT_NAME} Threads::Threads protobuf gRPC::grpc++)
