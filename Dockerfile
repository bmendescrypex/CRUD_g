# Usar a imagem oficial do Go
FROM golang:1.20-alpine

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Baixar as dependências do Go
RUN go mod download

# Copiar o resto do código da aplicação
COPY . .

# Compilar o código
RUN go build -o main .

# Expor a porta 8080 para acessar o serviço
EXPOSE 8080

# Comando de execução da aplicação
CMD ["./main"]
