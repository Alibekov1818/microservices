package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"sync"
	"time"
	"token-service/internal/data"
	"token-service/internal/db"
	pb "token-service/pb"
	"token-service/pkg/config"
	"token-service/pkg/jsonlog"
)

type Server struct {
	pb.UnimplementedTokenServiceServer
	Models data.Models
	Wg     sync.WaitGroup
	Logger *jsonlog.Logger
}

func (s *Server) GetToken(ctx context.Context, r *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {

	token, err := s.Models.Tokens.New(r.UserId, 24*time.Hour, data.ScopeAuthentication)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &pb.GetTokenResponse{Token: token.Plaintext}, nil
}

func Start() {

	conf := config.GetConfig()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := db.OpenDB()

	if err != nil {
		logger.PrintFatal(err, nil)
	}

	logger.PrintInfo("database connection pool established", nil)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))

	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	server := grpc.NewServer()

	reflection.Register(server)

	pb.RegisterTokenServiceServer(server, &Server{
		Logger: logger,
		Models: data.NewModels(db),
	})

	log.Printf("Server listening at %d", conf.Port)

	//if err = s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %s", err)
	//}
	go serveGoroutine(server, lis)
}

func serveGoroutine(s *grpc.Server, lis net.Listener) {
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
