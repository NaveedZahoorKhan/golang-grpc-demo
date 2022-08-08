package main

import (
	"context"
	"flag"
	"log"

	"github.com/NaveedZahoorKhan/golang-grpc-demo/pb"
	"github.com/NaveedZahoorKhan/golang-grpc-demo/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")

	flag.Parse()

	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	res, err := laptopClient.CreateLaptop(context.Background(), req)
	if err != nil {
		st, ok := status.FromError(err)

		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("laptop already exist")
		} else {
			log.Fatal(err)
		}
	}
	log.Printf("laptop created: %s", res.Id)

	findLaptopReq := &pb.FindLaptopRequest{
		Id: res.Id,
	}

	findLaptopRes, err1 := laptopClient.FindLaptop(context.Background(), findLaptopReq)

	if err1 != nil {
		log.Fatal(err1)
	}
	log.Printf("found laptop with id %s", findLaptopRes.GetLaptop())

}
