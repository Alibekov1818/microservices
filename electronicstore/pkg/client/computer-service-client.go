package client

import (
	"electronicstore/internal/jsonlog"
	pb "electronicstore/pb/computers-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GetComputersClient() pb.ComputersServiceClient {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelError)
	dial, err := grpc.Dial("localhost:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//dial, err := grpc.Dial("computers:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.PrintError(err, nil)
	}
	return pb.NewComputersServiceClient(dial)
}
