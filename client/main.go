// // // // package client
// // // package main

// // // import (
// // // 	"context"
// // // 	"flag"
// // // 	"fmt"
// // // 	"io"
// // // 	"log"
// // // 	"os"

// // // 	"github.com/tranvinhhien/learngrpc/pb"
// // // 	"google.golang.org/grpc"
// // // 	"google.golang.org/grpc/credentials/insecure"
// // // )

// // // func main() {
// // // 	serverAddress := flag.String("address", "", "the server address")
// // // 	flag.Parse()
// // // 	log.Printf("dial server  %d", *serverAddress) // tat bao mat
// // // 	conn, err := grpc.NewClient(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
// // // 	if err != nil {
// // // 		log.Fatal("cannot conn client %w", err)
// // // 		return
// // // 	}

// // // 	// laptopClient := pb.NewLaptopServiceClient(conn)
// // // 	// laptop := generate.RandomLaptop(1)[0]
// // // 	// req := &pb.CreateLaptopRequest{
// // // 	// 	Laptop: laptop,
// // // 	// }
// // // 	// res, err := laptopClient.CreateLaptop(context.Background(), req)
// // // 	// if err != nil {
// // // 	// 	log.Println(err.Error())
// // // 	// 	return
// // // 	// }

// // //		// log.Println("create laptop with id = %d", res.)
// // //	}
// // package main

// // import (
// // 	"context"
// // 	"flag"
// // 	"fmt"
// // 	"io"
// // 	"log"
// // 	"os"

// // 	"github.com/tranvinhhien/learngrpc/pb"
// // 	"google.golang.org/grpc"
// // 	"google.golang.org/grpc/credentials/insecure"
// // )

// // func main() {
// // 	serverAddress := flag.String("address", "", "the server address")
// // 	flag.Parse()
// // 	log.Printf("dial server %s", *serverAddress)

// // 	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
// // 	if err != nil {
// // 		log.Fatalf("cannot connect to server: %v", err)
// // 		return
// // 	}
// // 	defer conn.Close()

// // 	laptopClient := pb.NewLaptopServiceClient(conn)

// // 	req := &pb.SearchLaptopRequest{
// // 		Filter: &pb.Filter{
// // 			MaxPriceUsd: 1000,
// // 			MinCpuCores: 3,
// // 			MinCpuGhz:   100,
// // 			MinRam: &pb.Memory{
// // 				Value: 20,
// // 				Unit:  pb.Memory_GIGABYTE,
// // 			},
// // 		},
// // 	}

// // 	stream, err := laptopClient.SearchLaptop(context.Background(), req)
// // 	if err != nil {
// // 		log.Fatalf("Error while calling SearchLaptop: %v", err)
// // 		return
// // 	}

// // 	for {
// // 		res, err := stream.Recv()
// // 		if err == io.EOF {
// // 			// Kết thúc stream
// // 			break
// // 		}
// // 		if err != nil {
// // 			log.Fatalf("Error while receiving stream: %v", err)
// // 		}

// //			fmt.Printf("Received laptop: %v\n", res.GetLaptop())
// //			os.Stdout.Sync() // Đảm bảo việc in ra màn hình ngay lập tức
// //		}
// //	}
// package main

// import (
// 	"context"
// 	"flag"
// 	"log"
// 	"time"

// 	"github.com/tranvinhhien/learngrpc/generate"
// 	"github.com/tranvinhhien/learngrpc/pb"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// func main() {
// 	serverAddress := flag.String("address", "", "the server address")
// 	flag.Parse()
// 	log.Printf("dial server %s", *serverAddress)

// 	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("cannot connect to server: %v", err)
// 		return
// 	}
// 	defer conn.Close()

// 	laptopClient := pb.NewLaptopServiceClient(conn)
// 	stream, err := laptopClient.TotalPriceLaptops(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error while TotalPriceLaptops stream: %v", err)
// 	}
// 	listLapTop := generate.RandomLaptop(10)

// 	var totalClient float64
// 	for _, laptop := range listLapTop {
// 		log.Println("price laptop : ", laptop.PriceUsd)
// 		err := stream.Send(&pb.TotalPriceLaptopRequest{Laptop: laptop})
// 		totalClient += laptop.PriceUsd
// 		time.Sleep(time.Second)
// 		if err != nil {
// 			log.Fatal("err while TotalPriceLaptops send stream", err)
// 		}
// 	}
// 	resp, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Fatal("err while CloseAndRecv ", err)
// 		return
// 	}

//		log.Println("total price laptops is :", resp.TotalPrice)
//		log.Println("total price laptops in client is :", totalClient)
//	}
package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"github.com/tranvinhhien/learngrpc/generate"
	"github.com/tranvinhhien/learngrpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect to server: %v", err)
		return
	}
	defer conn.Close()

	laptopClient := pb.NewLaptopServiceClient(conn)
	stream, err := laptopClient.ChatMSLaptops(context.Background())
	if err != nil {
		log.Fatalf("Error while TotalPriceLaptops stream: %v", err)
	}
	wait := make(chan struct{})
	go func() {
		start := time.Now()
		for {
			elapsed := time.Since(start)
			if elapsed > 20*time.Second {
				log.Printf("end time chatting")
				break
			}
			err := stream.Send(&pb.ChatMsRequest{
				Ms: &pb.MsInfo{
					ToId: uint32(generate.RandomInt(1, 2)),
					Ms:   generate.RandomString(generate.RandomInt(5, 30)),
				},
			})
			if err != nil {
				log.Fatalf("err while send ms,", err)
				break
			}
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				log.Println("end chat form server")
				break
			}
			if err != nil {
				log.Fatalf("err while Recv ms,", err)
				break
			}
			log.Println("messasge from :%d  content is : %s", req.Ms.ToId, req.Ms.Ms)
		}
		close(wait)
	}()
	<-wait
}
