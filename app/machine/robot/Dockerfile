#build stage
FROM hominsu/grpc-alpine:alpine3.15-build as builder
# app dir
ARG APP_RELATIVE_PATH
# install: git(makefile use git to describe the version)
RUN apk --no-cache add git
# build app
COPY . /src
WORKDIR /src
RUN make grpc-cpp
WORKDIR /src/app/${APP_RELATIVE_PATH}
RUN make build

#final stage
FROM hominsu/grpc-alpine:alpine3.15-runtime
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
