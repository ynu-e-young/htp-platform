#build stage
FROM golang:alpine as builder
# app dir
ARG APP_RELATIVE_PATH
# install: gcc、g++、make、openssl、c-ares、cmake、grpc、protobuf、protoc、nlohmann-json
RUN apk --no-cache add bash findutils build-base openssl-dev openssl-libs-static re2-dev cmake c-ares-dev \
    grpc grpc-dev protobuf-dev libprotobuf protoc libprotoc nlohmann-json jsoncpp-dev jsoncpp-static
# build app
COPY . /src
WORKDIR /src
RUN make grpc-cpp
WORKDIR /src/app/${APP_RELATIVE_PATH}
RUN make build

#final stage
FROM alpine:latest
# image info
ARG AUTHOR_NAME
ARG AUTHOR_EMAIL
ARG VERSION
# label
LABEL author=$AUTHOR_NAME email=${AUTHOR_EMAIL} version=$VERSION
# app dir
ARG APP_RELATIVE_PATH
# install: tzdata c-ares libprotobuf
RUN apk --no-cache add tzdata c-ares libprotobuf
# cpoy the app from builder
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app
WORKDIR /app
EXPOSE 8000 9000
VOLUME /data/conf
CMD ["./server"]
