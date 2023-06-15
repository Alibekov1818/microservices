package main

import (
	"context"
	"electronicstore/internal/data"
	"electronicstore/internal/validator"
	pb "electronicstore/pb/token-service"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"net/http"
	"time"
)

func (app *application) createToken(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	v := validator.New()
	data.ValidateEmail(v, input.Email)
	data.ValidatePasswordPlaintext(v, input.Password)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.invalidCredentialsResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	match, err := user.Password.Matches(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if !match {
		app.invalidCredentialsResponse(w, r)
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
		"token_client", // name
		false,          // durable
		true,           // delete when unused
		false,          // exclusive
		false,          // noWait
		nil,            // arguments
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

	bytes, err := proto.Marshal(&pb.GetTokenRequest{
		UserId: user.ID,
	})
	if err != nil {
		app.logError(r, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",      // exchange
		"token", // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			ReplyTo:     q.Name,
			Body:        bytes,
		})
	if err != nil {
		app.logError(r, err)
	}

	for d := range msgs {
		res := &pb.GetTokenResponse{}

		err = proto.Unmarshal(d.Body, res)
		if err != nil {
			app.logError(r, err)
		}
		err = app.writeJSON(w, http.StatusCreated, envelope{"token": res.Token}, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
		break
	}
}
