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
	"phones_service/internal/data"
	"phones_service/internal/db"
	"phones_service/internal/validator"
	pb "phones_service/pb"
	"phones_service/pkg/config"
	"phones_service/pkg/jsonlog"
	"sync"
)

type Server struct {
	pb.UnimplementedPhonesServiceServer
	Models data.Models
	Wg     sync.WaitGroup
	Logger *jsonlog.Logger
}

func (s *Server) GetPhone(ctx context.Context, r *pb.PhoneId) (*pb.Phone, error) {
	phone, err := s.Models.Phones.Get(r.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Phone not found")
	}
	return &pb.Phone{Id: phone.ID, Model: phone.Model, Brand: phone.Brand, Year: phone.Year, Price: phone.Price}, nil
}

func (s *Server) GetPhones(context.Context, *pb.GetPhonesRequest) (*pb.PhoneList, error) {
	list, err := s.Models.Phones.SelectAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can not get list")
	}
	var phonesAll []*pb.Phone

	for i := 0; i < len(list); i++ {
		phonesAll = append(phonesAll, &pb.Phone{
			Id:    list[i].ID,
			Model: list[i].Model,
			Brand: list[i].Brand,
			Year:  list[i].Year,
			Price: list[i].Price,
		})
	}
	return &pb.PhoneList{List: phonesAll}, nil
}

func (s *Server) CreatePhone(ctx context.Context, r *pb.Phone) (*pb.Phone, error) {
	newPhone := &data.Phone{Model: r.Model, Brand: r.Brand, Year: r.Year, Price: r.Price}
	v := validator.New()
	if data.ValidatePhone(v, newPhone); !v.Valid() {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid data")
	}
	err := s.Models.Phones.Insert(newPhone)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return r, nil
}

func (s *Server) DeletePhone(ctx context.Context, r *pb.PhoneId) (*pb.Phone, error) {
	err := s.Models.Phones.Delete(r.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't delete phones or phone not found")
	}
	return &pb.Phone{Id: r.Id}, nil
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

	pb.RegisterPhonesServiceServer(server, &Server{
		Logger: logger,
		Models: data.NewModels(db),
	})

	log.Printf("Server listening at %d", conf.Port)

	//server.Serve(lis)
	go serveGoroutine(server, lis)
}

func serveGoroutine(s *grpc.Server, lis net.Listener) {
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
