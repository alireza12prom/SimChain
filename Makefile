build:
	go mod tidy
	go build -o bin ./cmd/simchain

run: build
	./bin

clean:
	rm -f bin