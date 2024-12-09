package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/tranvinhhien/learngrpc/pb"
	"github.com/tranvinhhien/learngrpc/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("part", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)
	laptopserver := service.NewLaptopServiceClient()
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopserver)
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server %d", port)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server %d", port)
	}

}
