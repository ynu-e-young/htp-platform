# build stage
FROM golang:alpine AS builder
# app dir
ARG APP_RELATIVE_PATH
# intstall git build-base opencv-dev
RUN apk add --no-cache git build-base opencv-dev
# build app
COPY . /src
WORKDIR /src/app/${APP_RELATIVE_PATH}
RUN make build

# final stage
FROM alpine:latest
# image info
ARG AUTHOR_NAME
ARG AUTHOR_EMAIL
ARG VERSION
# label
LABEL author=$AUTHOR_NAME email=${AUTHOR_EMAIL} version=$VERSION
# app dir
ARG APP_RELATIVE_PATH
# install: tzdata ca-certificates
RUN apk --no-cache add tzdata ca-certificates opencv-dev
# cpoy the app from builder
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app
WORKDIR /app
EXPOSE 8000 9000
VOLUME /data/conf
CMD ["./server", "-conf", "/data/conf"]
