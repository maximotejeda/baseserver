# must create a .env file with info
# must have compose installed
include .env
export
OS:=${shell go env GOOS}
ARCH=$(shell go env GOARCH)
OOSS="linux"
ARRCHS="arm 386"
DEBUG=1
SERVICE=server
# can be docker or podman or whatever
CONTAINERS=podman
.phony: all clean build test clean-image build-image build-image-debug run-image run-image-debug run-local

build-image: clean 
	@$(CONTAINERS)-compose -f ./docker-compose.yaml build

run-image: build-image
	@$(CONTAINERS)-compose -f docker-compose.yaml up

build-image-debug: clean
	@$(CONTAINERS)-compose -f docker-compose-debug.yaml build

run-image-debug: build-image-debug
	@$(CONTAINERS)-compose -f docker-compose-debug.yaml up

run-local:build
	@bin/$(SERVICE)-$(OS)-$(ARCH)
build:
	@go build -o ./bin/$(SERVICE)-$(OS)-$(ARCH) ./server.go
test:
	@go test ./$(SERVICE)/...
clean:
	@rm -rf ./bin

clean-image:
	@$(CONTAINERS) system prune -f


