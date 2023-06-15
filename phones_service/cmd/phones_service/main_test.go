package main

import (
	"context"
	_ "database/sql"
	_ "log"
	"os"
	"phones_service/internal/data"
	db "phones_service/internal/db"
	server "phones_service/internal/grpc_server"
	"phones_service/internal/validator"
	"phones_service/pb"
	"phones_service/pkg/jsonlog"
	"testing"
)

var testServer server.Server
var id int64
var size int

func init() {
	db, _ := db.OpenDB()
	testServer = server.Server{
		Logger: jsonlog.New(os.Stdout, jsonlog.LevelInfo),
		Models: data.NewModels(db),
	}

}

func TestValidatePhone(t *testing.T) {
	v := validator.New()

	testPhone := &data.Phone{ID: 10, Model: "Test", Brand: "Test", Year: 2000, Price: 10000}

	if data.ValidatePhone(v, testPhone); !v.Valid() {
		t.Errorf("Should be valid")
	}

	testPhone = &data.Phone{ID: 123, Model: "Test", Brand: "", Year: 0, Price: 0}

	if data.ValidatePhone(v, testPhone); v.Valid() {
		t.Errorf("Should be invalid")
	}
}

func TestInsertPhone(t *testing.T) {
	testPhone := &data.Phone{ID: 10, Model: "Test", Brand: "Test", Year: 2000, Price: 10000}
	err := testServer.Models.Phones.Insert(testPhone)
	id = testPhone.ID
	if err != nil {
		t.Errorf("%v", err.Error())
	}
}

func TestGetPhone(t *testing.T) {
	_, err := testServer.GetPhone(context.Background(), &pb.PhoneId{Id: id})
	if err != nil {
		t.Errorf("%v", err.Error())
	}
}

func TestGetPhonesFirst(t *testing.T) {
	list, err := testServer.GetPhones(context.Background(), &pb.GetPhonesRequest{})
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	size = len(list.List)
}

func TestDeletePhone(t *testing.T) {
	err := testServer.Models.Phones.Delete(id)
	if err != nil {
		t.Errorf("%v", err.Error())
	}
}

func TestGetPhonesSecond(t *testing.T) {
	list, err := testServer.GetPhones(context.Background(), &pb.GetPhonesRequest{})
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if len(list.List)+1 != size {
		t.Errorf("Get computers error")
	}
}
