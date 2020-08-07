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

.PHONY: run.db
run.db:
	docker run --rm \
		--name pg-docker \
		-e POSTGRES_PASSWORD=postgres \
		-p 127.0.0.1:5432:5432 \
		-v ~/docker/volumes/postgres:/var/lib/postgresql/data \
		-d postgres
	docker run --rm \
		--name pd-admin \
		-p 80:80 \
		-v ~/docker/volumes/pg-admin:/var/lib/pgadmin \
		-v ~/docker/volumes/pg-servers.json:/pgadmin4/servers.json \
		-e 'PGADMIN_DEFAULT_EMAIL=admin@gocart.com' \
		-e 'PGADMIN_DEFAULT_PASSWORD=admin' \
		-d dpage/pgadmin4

.PHONY: run
run: build
	godotenv ./dist/${APP_NAME}

.PHONY: test
test:
	go test ./... \
	-cover
