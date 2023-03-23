#build stage
FROM alpine:latest AS builder
# app dir
ARG APP_RELATIVE_PATH
# install: git(makefile use git to describe the version)
RUN apk --no-cache add bash findutils git build-base cmake \
    openssl-dev openssl-libs-static c-ares-dev re2-dev protobuf-dev grpc-dev
# build app
COPY . /src
WORKDIR /src
RUN cd /src/app/machine/robot
WORKDIR /src/app/${APP_RELATIVE_PATH}
RUN make build

#final stage
FROM alpine:latest
# image info
ARG AUTHOR_NAME
ARG AUTHOR_EMAIL
ARG VERSION
# label
LABEL author=${AUTHOR_NAME} email=${AUTHOR_EMAIL} version=${VERSION}
# app dir
ARG APP_RELATIVE_PATH
# install: tzdata c-ares libprotobuf
RUN apk --no-cache add tzdata ca-certificates c-ares re2 libprotobuf grpc-cpp
# cpoy the app from builder
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app
WORKDIR /app
EXPOSE 8000 9000
VOLUME /data/conf
CMD ["./server"]
