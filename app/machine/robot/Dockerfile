#build stage
FROM hominsu/grpc-alpine:grpc-1.46.3-alpine-3.16-devel as builder
# app dir
ARG APP_RELATIVE_PATH
# install: git(makefile use git to describe the version)
RUN apk --no-cache add bash findutils git build-base cmake \
    openssl-dev openssl-libs-static
# build app
COPY . /src
WORKDIR /src
RUN make grpc-cpp && cd /src/app/machine/robot && make conf
WORKDIR /src/app/${APP_RELATIVE_PATH}
RUN make build

#final stage
FROM hominsu/grpc-alpine:grpc-1.46.3-alpine-3.16-runtime
# image info
ARG AUTHOR_NAME
ARG AUTHOR_EMAIL
ARG VERSION
# label
LABEL author=$AUTHOR_NAME email=${AUTHOR_EMAIL} version=$VERSION
# app dir
ARG APP_RELATIVE_PATH
# install: tzdata c-ares libprotobuf
RUN apk --no-cache add tzdata ca-certificates
# cpoy the app from builder
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app
WORKDIR /app
EXPOSE 8000 9000
VOLUME /data/conf
CMD ["./server"]
