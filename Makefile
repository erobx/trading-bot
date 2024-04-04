run: build
	@./bin/trade.exe

build:
	@go build -o bin/trade.exe ./cmd/...