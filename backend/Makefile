# Variables
go_bin ?= go
OPENAPI=../openapi.yaml
GEN_PKG=internal/openapigen

# Default target
.DEFAULT_GOAL := help

help:
	@echo "Available targets:"
	@echo "  make dep        	- Install dependencies & vendor"
	@echo "  make run        	- Run the app locally (Pure Go mode)"
	@echo "  make test       	- Run unit tests"
	@echo "  make build      	- Build binary into ./bin/backend"

dep:
	@go mod tidy
	@go mod vendor

run:
	CGO_ENABLED=0 $(go_bin) run main.go

test:
	CGO_ENABLED=0 $(go_bin) test ./... -v

build:
	CGO_ENABLED=0 $(go_bin) build -o bin/backend main.go

# OpenAPI Generation
tool-openapi:
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

openapi-gen: tool-openapi
	@echo "generating openapi types/server (oapi-codegen v2)"
	@oapi-codegen -generate "types,chi-server,spec" -package openapigen -o $(GEN_PKG)/openapi.gen.go $(OPENAPI)
