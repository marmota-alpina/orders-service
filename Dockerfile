# Etapa 1: Construção do binário Go
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Instala dependências necessárias para a compilação
RUN apk add --no-cache git

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod tidy

# Copia o código da aplicação
COPY . .

# Compila a aplicação para produção
RUN go build -o orders-service ./cmd/main.go

# Etapa 2: Imagem final (Alpine)
FROM alpine:3.18

WORKDIR /app

# Instala as dependências necessárias (incluindo o migrate)
RUN apk add --no-cache ca-certificates postgresql-client curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate

# Copia o binário compilado e os arquivos de migrations
COPY --from=builder /app/orders-service /app/orders-service
COPY migrations /app/migrations

# Define as portas que a aplicação expõe
EXPOSE 8080 8081 8082

# Comando de entrada para rodar as migrations antes de iniciar a aplicação
ENTRYPOINT ["/bin/sh", "-c", "migrate -path /app/migrations -database postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable up && /app/orders-service"]
