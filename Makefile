.PHONY: test test-coverage fmt lint run clean

# Variáveis
APP_NAME=flux-control

# Comandos
test:
	@echo "Executando testes..."
	go test ./internal/... -v

test-coverage:
	@echo "Executando testes com cobertura..."
	go test ./internal/... -coverprofile=coverage.out
	go tool cover -html=coverage.out

fmt:
	@echo "Formatando o código..."
	go fmt ./...

lint:
	@echo "Executando o linter..."
	golangci-lint run ./...

run:
	@echo "Executando a aplicação..."
	go run cmd/app/main.go

clean:
	@echo "Limpando binários..."
	go clean
	rm -f coverage.out

help:
	@echo "Comandos disponíveis:"
	@echo "  make test          - Executa todos os testes"
	@echo "  make test-coverage - Executa testes com cobertura"
	@echo "  make fmt           - Formata o código"
	@echo "  make lint          - Executa o linter"
	@echo "  make run           - Executa a aplicação"
	@echo "  make clean         - Limpa os binários"
	@echo "  make help          - Exibe esta ajuda"

# Comando padrão
default: help 