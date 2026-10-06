STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo

install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/local/sentinelone/sentinelone.plugin -tags "${BUILD_TAGS}" *.go

test:
	go test ./...

.PHONY: install test
