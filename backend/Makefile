.PHONY: run build test clean

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

migrate-up:
	migrate -path migrations -database "mysql://root:@tcp(localhost:3306)/ai_education" up

migrate-down:
	migrate -path migrations -database "mysql://root:@tcp(localhost:3306)/ai_education" down