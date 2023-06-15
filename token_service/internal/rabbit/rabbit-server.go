package rabbit

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"log"
	pb "token-service/pb"
)

func StartRabbitServer() {
	connGrpc, err := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer connGrpc.Close()

	client := pb.NewTokenServiceClient(connGrpc)
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalln("Rabbit connection error")
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalln("Rabbit connection error")
	}
	//defer channel.Close()

	q, err := channel.QueueDeclare(
		"token", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalln("Queue declaration error")
	}

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Fatalln("Consumer declaration error")
	}

	var getMessages chan struct{}

	go func() {

		ctx := context.Background()

		for d := range msgs {
			req := &pb.GetTokenRequest{}

			proto.Unmarshal(d.Body, req)

			res, err := client.GetToken(ctx, req)

			responseByte, err := proto.Marshal(res)
			if err != nil {
				log.Fatalf("could not marshal: %v", err)
			}
			//
			err = channel.PublishWithContext(ctx,
				"",             // exchange
				"token_client", // routing key
				false,          // mandatory
				false,          // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        responseByte,
				})
			if err != nil {
				log.Fatalln("Publishing error")
			}
			d.Ack(false)
		}
	}()
	<-getMessages
}
