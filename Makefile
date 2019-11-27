GO_FILES?=$(shell find . -name '*.go')
BUILD_LDFLAGS = "-s -w"

.PHONY: default
default: build

.PHONY: build
build: bin/vultr-cloud-controller-manager

bin/vultr-cloud-controller-manager: $(GO_FILES)
	GOOS=`go env GOOS` GOARCH=`go env GOARCH` CGO_ENABLED=0 \
		go build \
		-ldflags $(BUILD_LDFLAGS) \
		-o bin/vultr-cloud-controller-manager \
		*.go

.PHONY: run
run:
	go run $(CURDIR)/*.go $(ARGS)

.PHONY: clean
clean:
	rm -Rf bin/*
