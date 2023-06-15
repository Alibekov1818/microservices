package main

import (
	server "phones_service/internal/grpc_server"
	"phones_service/internal/rabbit"
)

func main() {
	server.Start()
	rabbit.StartRabbitServer()
}
