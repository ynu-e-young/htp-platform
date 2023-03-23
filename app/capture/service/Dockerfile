# build stage
FROM gocv/opencv:4.7.0 AS builder
# app dir
ARG APP_RELATIVE_PATH
# build app
COPY . /src
WORKDIR /src/app/${APP_RELATIVE_PATH}
RUN make build

# final stage
FROM gocv/opencv:4.7.0
# image info
ARG AUTHOR_NAME
ARG AUTHOR_EMAIL
ARG VERSION
# label
LABEL author=${AUTHOR_NAME} email=${AUTHOR_EMAIL} version=${VERSION}
# app dir
ARG APP_RELATIVE_PATH
# install: ca-certificates netbase
RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y
# cpoy the app from builder
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app
WORKDIR /app
EXPOSE 8000 9000
VOLUME /data/conf
CMD ["./server", "-conf", "/data/conf"]
