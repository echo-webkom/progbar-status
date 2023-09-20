build:
	@go build -o bin/status

run: build
	./bin/status