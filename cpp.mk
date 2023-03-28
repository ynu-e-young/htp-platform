APP_VERSION			:=	$(shell git describe --tags --always)
APP_RELATIVE_PATH	:=	$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)

.PHONY: dep conf ent wire build clean run test

# generate config define code
conf:
	@buf generate --path src/conf --template src/conf/buf.conf.gen.yaml

# download dependencies of module
dep:
	@echo "cpp no support"

# generate ent code
ent:
	@echo "cpp no support"

# generate wire code
wire:
	@echo "cpp no support"

# build golang application
build:
	@mkdir -p bin/
	@cmake -H. -Bbuild -DCMAKE_BUILD_TYPE=Release
	@cmake --build ./build --parallel
	@mv build/htp-platform.machine.robot bin/server
	@rm -rf build

# clean build files
clean:
	@rm -rf bin/

# run application
run:
	@echo "cpp no support"

# run tests
test:
	@echo "cpp no support"

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