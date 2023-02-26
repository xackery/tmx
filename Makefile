PROTO_FILES=$(shell find pb -name '*.proto')
.PHONY: proto
proto:
	@protoc \
	-I. \
	$(PROTO_FILES) \
	--go_out=.
example-all:
	@make example NAME=desert
	@make example NAME=perspective_walls
	@make example NAME=sewers
	@make example NAME=hexagonal-mini
	@make example NAME=test_hexagonal_tile_60x60x30
	@make example NAME=island
	@make example NAME=sewers DIR=sewer_automap/
	@make example NAME=lttp
example:
	@go run main.go examples/$(DIR)$(NAME).tmx out/$(NAME).json
test:
	@go test ./...
PACKAGE_NAME=tmx
PACKAGE_VERSION=0.0.2
.PHONY: build-all
build-all:
	@mkdir -p bin
	@rm -rf bin/*
	@make build BUILD=windows EXT=.exe ARCH=amd64
	@make build BUILD=windows EXT=.exe ARCH=386
	@make build BUILD=linux ARCH=amd64
	@make build BUILD=linux ARCH=386
	@make build BUILD=darwin ARCH=amd64
	@make build BUILD=darwin ARCH=386
.PHONY: build
build:
	@GOARCH=$(ARCH) GOOS=$(BUILD) go build -o bin/$(PACKAGE_NAME)_$(PACKAGE_VERSION)_$(BUILD)_$(ARCH)$(EXT) -ldflags "main.version=$(PACKAGE_VERSION)" main.go
sanitize:
	@echo "sanitize: checking for errors"
	rm -rf vendor/
	go vet -tags ci ./...
	test -z $(goimports -e -d . | tee /dev/stderr)
	gocyclo -over 30 .
	golint -set_exit_status $(go list -tags ci ./...)
	staticcheck -go 1.14 ./...
	go test -tags ci -covermode=atomic -coverprofile=coverage.out ./...
    coverage=`go tool cover -func coverage.out | grep total | tr -s '\t' | cut -f 3 | grep -o '[^%]*'`
