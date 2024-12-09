package service_test

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tranvinhhien/learngrpc/generate"
	"github.com/tranvinhhien/learngrpc/pb"
	"github.com/tranvinhhien/learngrpc/service"
	"google.golang.org/grpc"
)

func TestSaveCreateLaptop(t *testing.T) {
	item := service.NewLaptopServiceClient()
	CreateLaptopResponse, err := item.CreateLaptop(context.Background(), &pb.CreateLaptopRequest{Laptop: generate.RandomLaptop(1)[0]})
	require.NoError(t, err)
	require.NotEmpty(t, CreateLaptopResponse)

}

// func TestClientCreateLaptop(t *testing.T) {
// 	t.Parallel()

// 	laptoServer, serverAddress := startTestLaptopServer(t)
// 	laptopClient := newTestLaptopClient(t)
// }

func startTestLaptopServer(t *testing.T) (*service.LaptopServiceClient, string) {
	laptopServer := service.NewLaptopServiceClient()
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	go grpcServer.Serve(listener)
	return laptopServer, listener.Addr().String()
}

// func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
// 	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
// }
