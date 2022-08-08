package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"example.com/grpc/pb"
	"example.com/grpc/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on") // port to listen on

	flag.Parse()

	log.Printf("Starting server on port %d", *port)

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal(err)
	}

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal(err)
	}

}
