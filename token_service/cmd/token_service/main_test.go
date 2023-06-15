package main

import (
	"context"
	_ "database/sql"
	_ "log"
	"os"
	"testing"
	"time"
	"token-service/internal/data"
	db "token-service/internal/db"
	server "token-service/internal/grpc_server"
	"token-service/internal/validator"
	"token-service/pb"
	"token-service/pkg/jsonlog"
)

var testServer server.Server
var id int64
var token *data.Token

func init() {
	db, _ := db.OpenDB()
	testServer = server.Server{
		Logger: jsonlog.New(os.Stdout, jsonlog.LevelInfo),
		Models: data.NewModels(db),
	}

}

func TestValidateTokenPlainText(t *testing.T) {
	v := validator.New()

	invalidToken := "blablabla"
	validToken := "JU6NDPI7ZRL7FCKRXTSCISKBFE"

	if data.ValidateTokenPlaintext(v, validToken); !v.Valid() {
		t.Errorf("Should be valid")
	}

	if data.ValidateTokenPlaintext(v, invalidToken); v.Valid() {
		t.Errorf("Should be invalid")
	}
}

func TestGenerateToken(t *testing.T) {
	newToken, err := data.GenerateToken(1, 24*time.Hour, data.ScopeAuthentication)
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if newToken.Plaintext == "" {
		t.Errorf("Should not return empty token")
	}
	token = newToken
}

func TestInsertToken(t *testing.T) {
	err := testServer.Models.Tokens.Insert(token)
	if err != nil {
		t.Errorf("%v", err.Error())
	}
}

func TestGetTokenFirst(t *testing.T) {
	_, err := testServer.GetToken(context.Background(), &pb.GetTokenRequest{UserId: 1})
	if err != nil {
		t.Errorf("%v", err.Error())
	}
}

func TestGetTokenSecond(t *testing.T) {
	_, err := testServer.GetToken(context.Background(), &pb.GetTokenRequest{UserId: 80000})
	if err == nil {
		t.Errorf("Should get an error")
	}
}
