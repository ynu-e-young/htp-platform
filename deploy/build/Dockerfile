# build stage
FROM golang:alpine AS builder
# app dir
ARG APP_RELATIVE_PATH
# intstall git build-base
RUN apk add --no-cache git build-base
# build app
COPY . /src
WORKDIR /src/app/${APP_RELATIVE_PATH}
RUN make build

#final stage
FROM alpine:latest
# app dir
ARG APP_RELATIVE_PATH
# install: tzdata ca-certificates
RUN apk --no-cache add tzdata ca-certificates
# cpoy the app from builder
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app
WORKDIR /app
EXPOSE 8000 9000
VOLUME /data/conf
CMD ["./server", "-conf", "/data/conf"]
