CMAKE_MINIMUM_REQUIRED(VERSION 2.8)
if (POLICY CMP0025)
    # detect Apple's Clang
    cmake_policy(SET CMP0025 NEW)
endif ()
if (POLICY CMP0054)
    cmake_policy(SET CMP0054 NEW)
endif ()

set(LIB_MAJOR_VERSION "0")
set(LIB_MINOR_VERSION "4")
set(LIB_PATCH_VERSION "5")
set(LIB_VERSION_STRING "${LIB_MAJOR_VERSION}.${LIB_MINOR_VERSION}.${LIB_PATCH_VERSION}")

if (CMAKE_VERSION VERSION_LESS 3.0)
    PROJECT(htp-platform.machine.robot CXX)
    add_definitions(-DPROJECT_NAME=\"${PROJECT_NAME}\")
else ()
    cmake_policy(SET CMP0048 NEW)
    PROJECT(htp-platform.machine.robot
            VERSION "${LIB_VERSION_STRING}"
            LANGUAGES C CXX
            HOMEPAGE_URL "https://github.com/hominsu/htp-platform")
    add_definitions(-DPROJECT_NAME=\"${PROJECT_NAME}\")
    add_definitions(-DHOMEPAGE_URL=\"${PROJECT_HOMEPAGE_URL}\")
endif ()

# compile in release with debug info mode by default
if (NOT CMAKE_BUILD_TYPE)
    set(CMAKE_BUILD_TYPE "Release" CACHE STRING "Choose the type of build, options are: Debug Release RelWithDebInfo MinSizeRel." FORCE)
endif ()

option(HTP_PLATFORM_ENABLE_INSTRUMENTATION_OPT "Build neujson with -march or -mcpu options" ON)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED TRUE)

find_program(CCACHE_FOUND ccache)
if (CCACHE_FOUND)
    set_property(GLOBAL PROPERTY RULE_LAUNCH_COMPILE ccache)
    set_property(GLOBAL PROPERTY RULE_LAUNCH_LINK ccache)
    if (CMAKE_CXX_COMPILER_ID MATCHES "Clang")
        set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Qunused-arguments -fcolor-diagnostics")
    endif ()
endif (CCACHE_FOUND)

if (CMAKE_CXX_COMPILER_ID STREQUAL "GNU")
    if (HTP_PLATFORM_ENABLE_INSTRUMENTATION_OPT AND NOT CMAKE_CROSSCOMPILING)
        if (CMAKE_SYSTEM_PROCESSOR STREQUAL "powerpc" OR CMAKE_SYSTEM_PROCESSOR STREQUAL "ppc64" OR CMAKE_SYSTEM_PROCESSOR STREQUAL "ppc64le")
            set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -mcpu=native")
        else()
            set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -march=native")
        endif ()
    endif ()
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra -Werror -Wno-type-limits")
    set(EXTRA_CXX_FLAGS -Weffc++ -Wswitch-default -Wfloat-equal -Wconversion -Wsign-conversion)
elseif (CMAKE_CXX_COMPILER_ID MATCHES "Clang")
    if (NOT CMAKE_CROSSCOMPILING)
        if (CMAKE_SYSTEM_PROCESSOR STREQUAL "powerpc" OR CMAKE_SYSTEM_PROCESSOR STREQUAL "ppc64" OR CMAKE_SYSTEM_PROCESSOR STREQUAL "ppc64le")
            set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -mcpu=native")
        elseif (CMAKE_SYSTEM_PROCESSOR STREQUAL "arm64" AND CMAKE_SYSTEM_NAME MATCHES "Darwin")
            set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -mcpu=apple-m1")
        else()
            set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -march=native")
        endif ()
    endif ()
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra -Werror -Wno-missing-field-initializers")
    set(EXTRA_CXX_FLAGS -Weffc++ -Wswitch-default -Wfloat-equal -Wconversion -Wimplicit-fallthrough)
elseif (CMAKE_CXX_COMPILER_ID STREQUAL "MSVC")
    add_definitions(-D_CRT_SECURE_NO_WARNINGS=1)
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} /EHsc")
    # Always compile with /WX
    if (CMAKE_CXX_FLAGS MATCHES "/WX-")
        string(REGEX REPLACE "/WX-" "/WX" CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")
    else ()
        set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} /WX")
    endif ()
elseif (CMAKE_CXX_COMPILER_ID MATCHES "XL")
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -qarch=auto")
endif ()

if (CMAKE_SYSTEM_NAME MATCHES "Darwin")
    link_directories(/opt/homebrew/lib)
    include_directories(/opt/homebrew/include)
    set(OPENSSL_ROOT_DIR /opt/homebrew/Cellar/openssl@1.1/1.1.1q)
    set(OPENSSL_INCLUDE_DIR /opt/homebrew/Cellar/openssl@1.1/1.1.1q/include)
elseif (CMAKE_SYSTEM_NAME MATCHES "Linux")
    link_directories(/usr/local/lib)
    include_directories(/usr/local/include)
else ()
    message(FATAL_ERROR "Platform ${CMAKE_SYSTEM} is not support for this project")
endif ()

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

add_executable(${PROJECT_NAME} ${ProjectSRC})
target_link_libraries(${PROJECT_NAME} Threads::Threads protobuf gRPC::grpc++)
