package service_test

import (
	"context"
	"net"
	"testing"

	"github.com/NaveedZahoorKhan/golang-grpc-demo/pb"
	"github.com/NaveedZahoorKhan/golang-grpc-demo/sample"
	"github.com/NaveedZahoorKhan/golang-grpc-demo/serializer"
	"github.com/NaveedZahoorKhan/golang-grpc-demo/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, serverAddress := startTestLaptopServer(t)

	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()

	expectedId := laptop.Id

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)

	require.Equal(t, expectedId, res.Id)

	other, err := laptopServer.Store.Find(res.Id)

	require.NoError(t, err)

	require.NotNil(t, other)

	requireSameLaptop(t, laptop, other)

}

func requireSameLaptop(t *testing.T, laptop *pb.Laptop, other *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop)

	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(other)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
func newTestLaptopClient(t *testing.T, serverAddr string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	require.NoError(t, err)

	return pb.NewLaptopServiceClient(conn)
}
func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0")

	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}
