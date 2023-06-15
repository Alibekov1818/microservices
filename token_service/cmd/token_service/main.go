package main

import (
	server "token-service/internal/grpc_server"
	"token-service/internal/rabbit"
)

func main() {
	server.Start()
	rabbit.StartRabbitServer()
}
