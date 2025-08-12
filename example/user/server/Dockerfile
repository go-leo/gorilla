FROM golang:1.20

LABEL org.opencontainers.image.authors="soyacen@outlook.com"

WORKDIR /app

COPY . .

RUN go mod init github.com/go-leo/gorilla/example/user/server

RUN go mod tidy

RUN go build -o user ./main.go

EXPOSE 8000

ENTRYPOINT [ "user" ]
