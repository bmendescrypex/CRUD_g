# Usar uma imagem oficial do Golang
FROM golang:1.21.0-alpine

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar os arquivos do projeto para o container
COPY . .

# Baixar as dependências
RUN go mod tidy

# Compilar a aplicação Go
RUN go build -o main .

# Expor a porta que a aplicação irá usar
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]
