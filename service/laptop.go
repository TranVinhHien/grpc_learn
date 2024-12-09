package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/tranvinhhien/learngrpc/generate"
	pb "github.com/tranvinhhien/learngrpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//	type LaptopServiceClient interface {
//		CreateLaptop(
//			ctx context.Context,
//			req *pb.CreateLaptopRequest,
//		) (*pb.CreateLaptopResponse, error)
//	}
type LaptopServiceClient struct {
	pb.UnimplementedLaptopServiceServer
}

func NewLaptopServiceClient() *LaptopServiceClient {
	return &LaptopServiceClient{}
}

func (server *LaptopServiceClient) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot generate a new laptop id")
	}
	laptop.Id = uuid.String()

	// save to dbs
	err = saveLaptop(laptop)
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot save  a new laptop "+err.Error())
	}
	return &pb.CreateLaptopResponse{Id: laptop.Id}, nil
}

func (server *LaptopServiceClient) SearchLaptop(req *pb.SearchLaptopRequest, steam pb.LaptopService_SearchLaptopServer) error {
	fmt.Println(req)

	laptops := generate.RandomLaptop(10)
	for _, laptop := range laptops {
		res := pb.SearchLaptopResponse{Laptop: laptop}
		err := steam.Send(&res)
		time.Sleep(5 * time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
func (server *LaptopServiceClient) TotalPriceLaptops(stream pb.LaptopService_TotalPriceLaptopsServer) error {
	// var laptops []*pb.Laptop
	var total float64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("total price is : ", total)
			return stream.SendAndClose(&pb.TotalPriceLaptopResponse{
				TotalPrice: total,
			})
		}
		if err != nil {
			log.Fatalf("err when Recv LaptopService_TotalPriceLaptopsServer", err)
		}
		//append(laptops, laptop)
		total += req.Laptop.GetPriceUsd()
	}
}
func (server *LaptopServiceClient) ChatMSLaptops(stream pb.LaptopService_ChatMSLaptopsServer) error {
	log.Println("start chatting")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("end chatting")
			return nil
		}
		if err != nil {
			log.Fatalf("err while Recv :", err)
			return err
		}
		err = stream.Send(&pb.ChatMsResponse{Ms: req.Ms})
		if err != nil {
			log.Fatalf("err while Send message :", err)
			return err
		}
	}
}
