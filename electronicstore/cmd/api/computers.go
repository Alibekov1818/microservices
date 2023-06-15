package main

import (
	"context"
	"electronicstore/internal/data"
	"electronicstore/internal/validator"
	pb "electronicstore/pb/computers-service"
	"electronicstore/pkg/client"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"net/http"
	"time"
)

func (app *application) addComputer(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Model  string `json:"model"`
		Cpu    string `json:"cpu"`
		Memory string `json:"memory"`
		Price  int64  `json:"price"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	computer := &data.Computer{
		Model:  input.Model,
		Cpu:    input.Cpu,
		Memory: input.Memory,
		Price:  input.Price,
	}
	v := validator.New()

	if data.ValidateComputer(v, computer); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	createComputer, err := client.GetComputersClient().CreateComputer(context.Background(),
		&pb.Computer{Model: computer.Model, Cpu: computer.Cpu, Memory: computer.Memory, Price: computer.Price})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/computers/%d", createComputer.Id))
	err = app.writeJSON(w, http.StatusCreated, createComputer, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) getComputer(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		app.logError(r, err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		app.logError(r, err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"computers_client", // name
		false,              // durable
		true,               // delete when unused
		false,              // exclusive
		false,              // noWait
		nil,                // arguments
	)
	if err != nil {
		app.logError(r, err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		app.logError(r, err)
	}

	bytes, err := proto.Marshal(&pb.ComputerId{
		Id: id,
	})
	if err != nil {
		app.logError(r, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",          // exchange
		"computers", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			ReplyTo:     q.Name,
			Body:        bytes,
		})
	if err != nil {
		app.logError(r, err)
	}

	for d := range msgs {
		res := &pb.Computer{}

		err = proto.Unmarshal(d.Body, res)
		if err != nil {
			app.logError(r, err)
		}
		err = app.writeJSON(w, http.StatusOK, res, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
		return
	}
}

func (app *application) deleteComputer(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	_, err = client.GetComputersClient().DeleteComputer(context.Background(), &pb.ComputerId{Id: id})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"response": "Computer deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getComputers(w http.ResponseWriter, r *http.Request) {
	computers, err := client.GetComputersClient().GetComputers(context.Background(), &pb.GetComputersRequest{})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, computers, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
