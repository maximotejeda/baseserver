# must create a .env file with info
include .env
export
OS:=${shell go env GOOS}
ARCH=$(shell go env GOARCH)
OOSS="linux"
ARRCHS="arm 386"
DEBUG=1
SERVICE1=auth
SERVICE2=entrance
# can be docker or podman or whatever
CONTAINERS=podman
.phony: all clean build test clean-image build-image build-image-debug run-image run-image-debug run-local


