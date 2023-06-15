package main

import (
	"context"
	"electronicstore/internal/data"
	"electronicstore/internal/validator"
	pb "electronicstore/pb/phones-service"
	"electronicstore/pkg/client"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"net/http"
	"time"
)

func (app *application) addPhone(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Model string `json:"model"`
		Brand string `json:"brand"`
		Year  int64  `json:"year"`
		Price int64  `json:"price"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	phone := &data.Phone{
		Model: input.Model,
		Brand: input.Brand,
		Year:  input.Year,
		Price: input.Price,
	}
	v := validator.New()

	if data.ValidatePhone(v, phone); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	createPhone, err := client.GetPhonesClient().CreatePhone(context.Background(),
		&pb.Phone{Model: phone.Model, Brand: phone.Brand, Year: phone.Year, Price: phone.Price})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/phones/%d", createPhone.Id))
	err = app.writeJSON(w, http.StatusCreated, createPhone, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) getPhone(w http.ResponseWriter, r *http.Request) {
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
		"phones_client", // name
		false,           // durable
		true,            // delete when unused
		false,           // exclusive
		false,           // noWait
		nil,             // arguments
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

	bytes, err := proto.Marshal(&pb.PhoneId{
		Id: id,
	})
	if err != nil {
		app.logError(r, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",       // exchange
		"phones", // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			ReplyTo:     q.Name,
			Body:        bytes,
		})
	if err != nil {
		app.logError(r, err)
	}

	for d := range msgs {
		res := &pb.Phone{}

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

func (app *application) deletePhone(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	_, err = client.GetPhonesClient().DeletePhone(context.Background(), &pb.PhoneId{Id: id})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"response": "Phone deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getPhones(w http.ResponseWriter, r *http.Request) {
	phones, err := client.GetPhonesClient().GetPhones(context.Background(), &pb.GetPhonesRequest{})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, phones, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
