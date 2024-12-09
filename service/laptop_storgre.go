package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/tranvinhhien/learngrpc/pb"
)

func readData() []*pb.Laptop {
	// Đọc dữ liệu hiện có từ tệp JSON
	var laptops []*pb.Laptop

	// Kiểm tra xem tệp có tồn tại không
	if _, err := os.Stat("latop.json"); err == nil {
		file, err := ioutil.ReadFile("latop.json")
		if err != nil {
			return nil
		}

		// Giải mã JSON thành slice của Person
		err = json.Unmarshal(file, &laptops)
		if err != nil {
			return nil
		}
	}
	return laptops
}

func saveLaptop(laptop *pb.Laptop) error {
	// Đọc dữ liệu hiện có từ tệp JSON
	laptops := readData()

	// Thêm dữ liệu mới vào slice
	laptops = append(laptops, laptop)

	// Mã hóa lại slice thành JSON
	jsonData, err := json.MarshalIndent(laptops, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("latop.json", jsonData, 0644)
	if err != nil {
		return err
	}

	log.Println("Dữ liệu mới đã được thêm vào du_lieu.json")
	return nil //fmt.Errorf("id da bi trung lap")
}

// func Find(id string) (*pb.Laptop, error) {

// 	store.mutex.RLock()
// 	defer store.mutex.RUnlock()

// 	laptop := store.data[id]
// 	if laptop == nil {
// 		return nil, nil
// 	}

// 	return deepCopy(laptop)
// }

// Search searches for laptops with filter, returns one by one via the found function
func Search(
	ctx context.Context,
	filter *pb.Filter,
	found func(laptop *pb.Laptop) error,
) error {
	laptops := readData()

	for _, laptop := range laptops {
		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Print("context is cancelled")
			return nil
		}
		if isQualified(filter, laptop) {
			err := found(laptop)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func isQualified(filter *pb.Filter, laptop *pb.Laptop) bool {
	if laptop.GetPriceUsd() > filter.MaxPriceUsd {
		return false
	}
	if laptop.Cpu.GetNumbersCores() < filter.MinCpuCores {
		return false
	}
	if laptop.GetCpu().GetMinGhz() < filter.MinCpuGhz {
		return false
	}
	if toBit(laptop.GetRam()) < toBit(filter.MinRam) {
		return false
	}
	return true
}
func toBit(memory *pb.Memory) uint64 {
	value := memory.GetValue()
	switch memory.GetUnit() {
	case pb.Memory_BIT:
		return value
	case pb.Memory_BYTE:
		return value << 3
	case pb.Memory_KILOBYTE:
		return value << 13
	case pb.Memory_MEGABYTE:
		return value << 23
	case pb.Memory_GIGABYTE:
		return value << 33
	case pb.Memory_TERABYTE:
		return value << 43
	default:
		return 0
	}
}

// func deepCopy(laptop *pb.Laptop) (*pb.Laptop, error) {
// 	other := &pb.Laptop{}

// 	err := copier.Copy(other, laptop)
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
// 	}

// 	return other, nil
// }
