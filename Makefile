build:
	go build -o bin/gendiff ./cmd/gendiff
test:
	go test -v
clean:
	rm -rf bin/
lint:
	golangci-lint run