.PHONY: build
build:
	sqlc vet
	sqlc compile
	sqlc generate
varify:
	sqlc verify
run: build
	go run cmd/main.go
	