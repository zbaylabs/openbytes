# Variables
SERVICE=openbytes
PROTOC=PATH=${shell go env GOPATH}/bin:$(PATH) ${shell which protoc}
# Version information
#VERSION=1.0.0
REVISION=${shell git rev-parse --short HEAD}
#RELEASE=production
BUILD_HASH=${shell git rev-parse HEAD}
BUILD_TIME=${shell date "+%Y-%m-%d@%H:%M:%SZ%z"}
LD_FLAGS:=-X main.Version=$(VERSION) -X main.Revision=$(REVISION) -X main.Release=$(RELEASE) -X main.BuildHash=$(BUILD_HASH) -X main.BuildTime=$(BUILD_TIME)

prepare:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate: generate-go generate-js

generate-go:
	@$(PROTOC) -I./api --go_out=. --go-grpc_out=. api/*.proto
	@echo Generate-go completely.

generate-js:
	@$(PROTOC) -I./api api/*.proto \
	--js_out=import_style=commonjs:api/js \
	--grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:api/js
	@echo Generate-js completely.

ui:generate
#	cp -rf api/js/ ui/src/sdk
	cd ui && npm run build

build:ui
	rm -rf cmd/dist && cp -rf ui/dist cmd/ && rm -rf cmd/dist/js/*.map
	go build -ldflags='$(LD_FLAGS)' -o bundles/$(SERVICE) *.go