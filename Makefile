run-online: build-online
	@./bin/trade

build-online:
	@go build -o bin/trade ./cmd

templ:
	@templ generate --watch --proxy="http://localhost:3000" --open-browser=false -v

server:
	@air

tailwind:
	@npx tailwindcss -i pkg/view/css/input.css -o public/styles.css --watch

live:
	make -j3 tailwind templ server
