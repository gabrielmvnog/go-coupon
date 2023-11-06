FROM golang:1.21.3-alpine

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY src/ ./src
COPY proto/ ./proto

RUN go build src/main.go

CMD ["./main"]
