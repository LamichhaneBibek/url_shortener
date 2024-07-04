server:
	go run cmd/main.go

test:
	go test -v -cover ./...

.PHONY: server test