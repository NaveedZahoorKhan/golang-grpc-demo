package service

import (
	"context"
	"errors"
	"log"

	"github.com/NaveedZahoorKhan/golang-grpc-demo/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	Store LaptopStore
}

// FindLaptop implements pb.LaptopServiceServer
func (server *LaptopServer) FindLaptop(ctx context.Context, req *pb.FindLaptopRequest) (*pb.FindLaptopResponse, error) {
	laptopId := req.Id
	log.Printf("Received request to find laptop with id: %s", laptopId)

	if len(laptopId) > 0 {
		_, err := uuid.Parse(laptopId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	}
	laptop, err := server.Store.Find(laptopId)

	if err != nil {
		return nil, err
	}

	resp := &pb.FindLaptopResponse{
		Laptop: laptop,
	}
	return resp, nil
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{
		Store: store,
	}
}

func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("Received request to create laptop: %v", laptop)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		laptop.Id = id.String()
	}

	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExist) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "failed to save laptop")
	}
	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil

}
