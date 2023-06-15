package client

import (
	"electronicstore/internal/jsonlog"
	pb "electronicstore/pb/phones-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GetPhonesClient() pb.PhonesServiceClient {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelError)
	dial, err := grpc.Dial("localhost:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//dial, err := grpc.Dial("phones:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.PrintError(err, nil)
	}
	return pb.NewPhonesServiceClient(dial)
}
