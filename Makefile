VERSION := $(shell git describe --tags --always 2>/dev/null || echo "unknown")
BINARY := pinger-$(VERSION)
BUILD_DIR := build

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

config.yaml:
	cp ./template/config.yaml $@

.PHONY: build-android
build-android: |$(BUILD_DIR)
	GOOS=android GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY)-arm64 .

.PHONY: build-PC
build-PC: | $(BUILD_DIR)
	go build -o build/$(BINARY)-amd64 .

.PHONY: build
build: config.yaml build-android build-PC

.PHONY: clean
clean:
	rm -f $(BUILD_DIR)/pinger-*
	rmdir $(BUILD_DIR)