# imagem base
FROM golang:1.26-alpine

# pasta de trabalho
WORKDIR /app

# copiar arquivos do módulo
COPY go.mod go.sum ./

# instalar dependências
RUN go mod download

# copiar projeto
COPY . .

# build da aplicação
RUN go build -o app cmd/api/main.go

# expor porta
EXPOSE 8080

# rodar app
CMD ["./app"]