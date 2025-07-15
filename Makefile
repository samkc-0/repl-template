.DEFAULT_GOAL := help

BIN_DIR := bin
NAME ?= repl
BINARY := $(BIN_DIR)/$(shell echo $(NAME) | tr '[:upper:]' '[:lower:]')

.PHONY: help build clean config

# catch-all: e.g. `make pokedex`
%:
	@$(MAKE) build NAME=$@

build:
	@echo "generating config for: $(NAME)"
	@go run gen/genconfig.go $(NAME)
	@echo "building binary: $(BINARY)"
	@mkdir -p $(BIN_DIR)
	@go build -o $(BINARY)
	@echo "done"

config:
	@echo "generating config for: $(NAME)"
	@go run gen/genconfig.go $(NAME)
	@echo "config.go ready for 'go run .'"

clean:
	@echo "cleaning up..."
	@rm -f config.go
	@rm -rf $(BIN_DIR)
	@echo "done"

help:
	@echo ""
	@echo "usage:"
	@echo "  make <name>        build binary for given repl name (e.g. make pokedex)"
	@echo "  make config NAME=… generate config.go only (for go run .)"
	@echo "  make clean         remove binaries + generated config"
	@echo ""
