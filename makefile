.PHONY: all run test

run:
	go run cmd/server/main.go

stop:
	@echo "Parando processos na porta 8000..."
	-fuser -k 8000/tcp || true

test:
	go test -v ./...