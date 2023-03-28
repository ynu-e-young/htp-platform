GOPATH				:=	$(shell go env GOPATH)
APP_VERSION			:=	$(shell git describe --tags --always)
APP_RELATIVE_PATH	:=	$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)

.PHONY: dep conf ent wire build clean run test

# generate config define code
conf:
	@buf generate --path internal/conf --template internal/conf/buf.conf.gen.yaml

# download dependencies of module
dep:
	@go mod download

# generate ent code
ent:
	@if [ -d "./internal/data/ent" ]; then \
  		go run entgo.io/ent/cmd/ent generate \
				--feature privacy \
				--feature entql \
				--feature sql/modifier \
				--feature sql/execquery \
				--feature sql/upsert \
				./internal/data/ent/schema; \
	fi

# generate wire code
wire:
	@go run github.com/google/wire/cmd/wire ./cmd/server

# build golang application
build:
	@if [ ! -d "./bin/" ]; then mkdir bin; fi
	@go build -ldflags "-X main.Version=$(APP_VERSION)" -o ./bin/ ./...

# clean build files
clean:
	@go clean

# run application
run:
	@go run ./cmd/server -conf ./configs

# run tests
test:
	@go test -v ./... -cover

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help