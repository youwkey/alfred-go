# Dependency tool versions
GOTESTSUM_VERSION := 1.7.0
GOLANGCLI_VERSION := 1.44.2

BUILD_DIR := build

.PHONY: check test lint check-goreleaser clean tools

check: test lint check-goreleaser

test:
	@mkdir -p $(BUILD_DIR)
	gotestsum --format=short-verbose -- -race -coverprofile=$(BUILD_DIR)/coverage.txt -covermode=atomic
	go tool cover -html=$(BUILD_DIR)/coverage.txt -o $(BUILD_DIR)/coverage.html

lint:
	golangci-lint run

check-goreleaser:
	goreleaser check

clean:
	rm -rf $(BUILD_DIR)

tools:
	go install gotest.tools/gotestsum@v$(GOTESTSUM_VERSION)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCLI_VERSION)
