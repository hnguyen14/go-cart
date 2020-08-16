SRCROOT ?= $(realpath .)

APP_NAME ?= go-cart
APP_PORT ?= 8080
ENV_FILE ?= .env
MAJOR_VERSION = 1
MINOR_VERSION = 0
BUILD_NUMBER ?= 0
REVISION ?= $(shell git rev-parse --short HEAD)
VERSION = $(MAJOR_VERSION).$(MINOR_VERSION).$(BUILD_NUMBER)

.PHONY: build
build: vendor
		CGO_ENABLED=0 go build \
		-mod=vendor \
		-o dist/${APP_NAME} \
		-ldflags "-s -X main.version=$(VERSION) -X main.revision=$(REVISION)"

.PHONY: vendor
vendor: $(SRCROOT)/go.mod $(SRCROOT)/go.sum
		go mod vendor

.PHONY: run
run: build
	godotenv ./dist/$(APP_NAME)

.PHONY: test
test:
	go test ./... \
	-cover

.PHONY: docker.build
docker.build:
	docker build -t $(APP_NAME) .

.PHONY: docker.run
docker.run:
	docker run \
	--env-file $(ENV_FILE) \
	-p $(APP_PORT):8080 \
	-w /app \
	-d $(APP_NAME)
