build:
	@go build -o ./bin/web/snippetbox ./cmd/web

run: build
	@./bin/web/snippetbox