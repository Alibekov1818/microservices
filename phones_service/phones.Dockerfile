FROM golang:1.19-alpine

WORKDIR /app/phones_service
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/phones_service/main.go"]