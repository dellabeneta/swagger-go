FROM golang:1.23-alpine AS build

# Define o diretório de trabalho
WORKDIR /app

# Instala o Swag CLI para gerar documentação Swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copia os arquivos de código fonte
COPY . .

# Baixa as dependências
RUN go mod download

# Gera a documentação Swagger
RUN swag init

# Compila o aplicativo
# CGO_ENABLED=0 desativa o CGO para gerar um binário estático
# -ldflags="-s -w" remove informações de depuração para reduzir o tamanho
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o api-server .

# Segunda etapa - imagem mínima
FROM alpine:latest

# Adiciona certificados CA para possibilitar conexões HTTPS
RUN apk --no-cache add ca-certificates

# Cria um usuário não-root para executar o aplicativo
RUN adduser -D -H -h /app appuser

WORKDIR /app

# Copia o binário compilado da etapa anterior
COPY --from=build /app/api-server .

# Define o usuário não-root
USER appuser

# Expõe a porta utilizada pelo aplicativo
EXPOSE 8080

# Comando para executar o aplicativo
CMD ["./api-server"]