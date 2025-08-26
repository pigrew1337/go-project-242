lint:
	~/go/bin/golangci-lint run

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

go run:
	go run  cmd/hexlet-path-size/main.go