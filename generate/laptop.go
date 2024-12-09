package generate

import (
	"fmt"
	"sync"

	"github.com/tranvinhhien/learngrpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func RandomLaptop(length int) []*pb.Laptop {
	var listLaptop []*pb.Laptop
	var wg sync.WaitGroup

	for id := range length {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			listLaptop = append(listLaptop, &pb.Laptop{
				Id:          fmt.Sprintf("laptop-%d", id),
				Brand:       RandomString(10 + id),
				Name:        RandomString(10 + id),
				Cpu:         RandomCpu(),
				Ram:         RandomMemory(),
				Gpu:         RandomGPU(),
				Storegres:   RandomStorage(),
				Screen:      RandomScreen(),
				Keyboard:    RandomKeyboard(),
				Weight:      float64(RandomInt(4, 10)),
				PriceUsd:    float64(RandomInt(1000, 10000)),
				ReleaseYear: uint32(RandomInt(2020, 2024)),
				UpdateAt:    timestamppb.Now(),
			})
		}(id)
	}
	wg.Wait()
	return listLaptop
}

func RandomCpu() *pb.CPU {
	return &pb.CPU{
		Brand:        RandomString(10),
		Name:         RandomString(8),
		NumbersCores: uint32(RandomInt(1, 10)),
		MinGhz:       float64(RandomInt(1, 10)),
		MaxGhz:       float64(RandomInt(1, 10)),
	}
}

func RandomMemory() *pb.Memory {
	return &pb.Memory{
		Value: uint64(RandomInt(1, 10)),
		Unit:  pb.Memory_Unit(RandomInt(0, 6)),
	}
}
func RandomScreen() *pb.Screen {
	return &pb.Screen{
		SizeInch:      float32(RandomInt(5, 10)),
		Resolution:    &pb.Screen_Resolution{Width: uint32(RandomInt(50, 100)), Height: uint32(RandomInt(50, 100))},
		Panel:         pb.Screen_Panel(RandomInt(0, 2)),
		Multipletouch: RandomBool(),
	}
}
func RandomKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  pb.Keyboard_Layout(RandomInt(0, 3)),
		Backlit: RandomBool(),
	}
}
func RandomGPU() []*pb.GPU {
	var listGPU []*pb.GPU
	for i := range RandomInt(2, 5) {
		listGPU = append(listGPU, &pb.GPU{
			Brand:  RandomString(10 + i),
			Name:   RandomString(10 + i),
			MinGhz: float64(RandomInt(10, 100)),
			MaxGhz: float64(RandomInt(100, 500)),
			Memory: &pb.Memory{
				Value: uint64(RandomInt(10, 100)),
				Unit:  pb.Memory_Unit(RandomInt(0, 6)),
			},
		})
	}
	return listGPU
}

func RandomStorage() []*pb.Storage {
	var listStorage []*pb.Storage
	for _ = range RandomInt(2, 5) {
		listStorage = append(listStorage, &pb.Storage{
			Driver: pb.Storage_Driver(RandomInt(0, 2)),
			Memory: &pb.Memory{
				Value: uint64(RandomInt(10, 100)),
				Unit:  pb.Memory_Unit(RandomInt(0, 6)),
			},
		})
	}
	return listStorage
}
