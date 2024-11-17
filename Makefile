MODULE_NAME = $(shell go list -m)
BIN_DIR = bin
BINARY_NAME = $(notdir $(MODULE_NAME))

.PHONY: all
all: build

.PHONY: build
build: 
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(BINARY_NAME) .

.PHONY: run
run: build
	@go run .

.PHONY: test
test:
	@go test  ./...

.PHONY: clean
clean:
	@rm -rf $(BIN_DIR)