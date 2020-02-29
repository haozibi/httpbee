APP?=httpbee

# SHELL := /bin/bash # Use bash syntax
GOOS?=linux
GOARCH?=amd64

VERSION?=$(shell git describe --tags --always)
COMMIT_HASH?=$(shell git rev-parse --short HEAD 2>/dev/null)
NOW?=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')
PROJECT?=github.com/haozibi/${APP}

LDFLAGS += -X "${PROJECT}/app.BuildTime=${NOW}"
LDFLAGS += -X "${PROJECT}/app.BuildVersion=${VERSION}"
LDFLAGS += -X "${PROJECT}/app.BuildAppName=${APP}"
LDFLAGS += -X "${PROJECT}/app.CommitHash=${COMMIT_HASH}"
BUILD_TAGS = ""
BUILD_FLAGS = "-v"

.PHONY: build build-linux clean govet

default: build

build: clean govet
	CGO_ENABLED=0 GOOS= GOARCH= go build ${BUILD_FLAGS} -ldflags '${LDFLAGS}' -tags '${BUILD_TAGS}' -o ${APP} 


build-linux: clean govet
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build ${BUILD_FLAGS} -ldflags '${LDFLAGS}' -tags '${BUILD_TAGS}' -o ${APP}

build-all: clean
	rm -rf release && chmod +x build.sh && ./build.sh

govet: 
	@ go vet . && go fmt ./... 
	# && \
	# (if [[ "$(gofmt -d $(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./tests/*" -not -path "./assets/*"))" == "" ]]; then echo "Good format"; else echo "Bad format"; exit 33; fi);

clean: 
	@ rm -fr ${APP} main
