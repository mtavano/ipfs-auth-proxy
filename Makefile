#
# INTERNAL VARIABLES
#
BIN=$(PWD)/bin/
#
# TARGETS
#


dev:
	@echo "[dev] Running service in debug-hot-reload mode..."
	@export $$(cat .env) && nodemon --exec go run cmd/server/main.go --signal SIGTERM

build-arm:
	@echo "[build] Building arm service..."
	@GOOS=linux GOARCH=arm64 go build -o bin/ipfs-proxy-arm cmd/server/main.go

build-linux:
	@echo "[build] Building linux service..."
	@GOOS=linux go build -o bin/ipfs-proxy-linux cmd/server/main.go

build-macos:
	@echo "[build] Building macos service..."
	@GOOS=darwin go build -o bin/ipfs-proxy-macos cmd/server/main.go

.PHONY: build-arm
