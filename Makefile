run:
	go run cmd/main.go

test:
	go test ./...

build:
	go build -o go-api cmd/main.go