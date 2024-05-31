
GO := go


BUILD_FLAGS :=


SERVER_PATH := cmd/server
CLIENT_PATH := cmd/client
BIN_DIR := bin


.PHONY: all
all: server client

.PHONY: server
server:
	$(GO) build $(BUILD_FLAGS) -o $(BIN_DIR)/server $(SERVER_PATH)/main.go


.PHONY: client
client:
	$(GO) build $(BUILD_FLAGS) -o $(BIN_DIR)/client $(CLIENT_PATH)/main.go

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)
