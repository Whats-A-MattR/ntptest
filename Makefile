.PHONY: all local clean linux windows darwin

APP_NAME := ntptest
BIN_DIR := bin

all: linux windows darwin
	@echo "✅ Built all targets."

local:
	@echo "🔨 Building for local system..."
	go build -o $(APP_NAME) main.go
	@echo "✅ Done: ./$(APP_NAME)"

linux:
	@echo "🐧 Building Linux binary..."
	mkdir -p $(BIN_DIR)/linux
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/linux/$(APP_NAME) main.go
	@echo "✅ Done: $(BIN_DIR)/linux/$(APP_NAME)"

windows:
	@echo "🪟 Building Windows binary..."
	mkdir -p $(BIN_DIR)/windows
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/windows/$(APP_NAME).exe main.go
	@echo "✅ Done: $(BIN_DIR)/windows/$(APP_NAME).exe"

darwin:
	@echo "🍎 Building macOS binary..."
	mkdir -p $(BIN_DIR)/darwin
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/darwin/$(APP_NAME) main.go
	@echo "✅ Done: $(BIN_DIR)/darwin/$(APP_NAME)"

clean:
	@echo "🧹 Cleaning up..."
	rm -f $(APP_NAME)
	rm -rf $(BIN_DIR)
	@echo "✅ Cleaned."
