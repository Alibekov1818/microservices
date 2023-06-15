package client

import (
	"electronicstore/internal/jsonlog"
	pb "electronicstore/pb/token-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GetAuthClient() pb.TokenServiceClient {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelError)
	dial, err := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//dial, err := grpc.Dial("token:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.PrintError(err, nil)
	}
	return pb.NewTokenServiceClient(dial)
}
