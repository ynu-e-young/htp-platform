.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.6.1

.PHONY: api
# generate api
api:
	find app -mindepth 2 -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'

.PHONY: wire
# generate wire
wire:
	find app -mindepth 2 -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) wire'

.PHONY: proto
# generate proto
proto:
	find app -mindepth 2 -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) proto'

.PHONY: generate
# generate generate
generate:
	find app -mindepth 2 -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) generate'

.PHONY: build
# generate build
build:
	find app -mindepth 2 -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) build'

.PHONY: docker
# generate docker
docker:
	find app -mindepth 2 -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) docker'

.PHONY: buildx
# generate buildx
buildx:
	find app -mindepth 2 -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) buildx'