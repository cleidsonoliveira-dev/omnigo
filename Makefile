# Nome do binário (pega do módulo Go automaticamente)
APP_NAME := $(shell basename $(shell go list -m))

# Diretório de saída
BIN_DIR := bin

# Plataformas alvo
PLATFORMS := \
	linux/amd64 \
	linux/arm64 \
	darwin/amd64 \
	darwin/arm64 \
	windows/amd64

# Build default (para o SO atual)
build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) main.go

# Build para todas as plataformas
build-all:
	@mkdir -p $(BIN_DIR)
	@for platform in $(PLATFORMS); do \
		GOOS=$$(echo $$platform | cut -d'/' -f1); \
		GOARCH=$$(echo $$platform | cut -d'/' -f2); \
		output="$(BIN_DIR)/$(APP_NAME)-$$GOOS-$$GOARCH"; \
		if [ "$$GOOS" = "windows" ]; then output="$$output.exe"; fi; \
		echo ">> Building $$output"; \
		GOOS=$$GOOS GOARCH=$$GOARCH go build -o $$output main.go || exit 1; \
	done

# Testes unitários
test:
	go test -v ./...

# Testes de integração (marcados com a build tag "integration")
test-integration:
	go test -v -tags=integration ./...

# Executar localmente
run:
	go run main.go

# Linter (necessário instalar golangci-lint: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
lint:
	golangci-lint run ./...

# Cobertura de testes
cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Limpar binários e arquivos de cobertura
clean:
	rm -rf $(BIN_DIR) coverage.out

# Targets que não correspondem a arquivos
.PHONY: build build-all test test-integration run lint cover clean