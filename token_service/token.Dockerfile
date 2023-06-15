FROM golang:1.19-alpine

WORKDIR /app/token_service
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/token_service/main.go"]