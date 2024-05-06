run-online: build-online
	@./bin/trade.exe

run-local: build-local
	@./bin/sim.exe

build-online:
	@go build -o bin/trade.exe ./cmd/online/

build-local:
	@go build -o bin/sim.exe ./cmd/local/

test: build-test
	@./bin/test-sim.exe

build-test:
	@go build -o bin/test-sim.exe ./test/cmd_sim/