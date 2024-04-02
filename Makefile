run: build
	@./bin/trade.exe

build:
	@go build -o bin/trade.exe cmd/main.go cmd/simulation.go cmd/skin.go cmd/user.go cmd/market.go