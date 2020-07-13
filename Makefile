SRCROOT ?= $(realpath .)

APP_NAME ?= go-cart
MAJOR_VERSION = 1
MINOR_VERSION = 0
BUILD_NUMBER ?= 0
REVISION ?= $(shell git rev-parse --short HEAD)
VERSION = $(MAJOR_VERSION).$(MINOR_VERSION).$(BUILD_NUMBER)

vendor: $(SRCROOT)/go.mod $(SRCROOT)/go.sum
		go mod vendor

.PHONY: build
build: vendor
		go build \
		-mod=vendor \
		-o dist/${APP_NAME} \
		-ldflags "-s -X main.version=$(VERSION) -X main.revision=$(REVISION)"

.PHONY: run
run: build
	godotenv ./dist/${APP_NAME}
.PHONY: test
test:
	go test ./... \
	-cover