package sample

import (
	"time"

	"example.com/grpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        float32(minGhz),
		MaxGhz:        maxGhz,
	}

	return cpu
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_UNIT_GIGABYTE,
	}
	gpu := &pb.GPU{
		Name:   name,
		Brand:  brand,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu
}

func NewRam() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  pb.Memory_UNIT_GIGABYTE,
	}
	return ram
}
func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb.Memory_UNIT_GIGABYTE,
		},
	}
	return ssd
}

func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_UNIT_TERABYTE,
		},
	}
	return hdd
}
func NewScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:          randomID(),
		Brand:       brand,
		Name:        name,
		Cpu:         NewCPU(),
		Ram:         NewRam(),
		Gpus:        []*pb.GPU{NewGPU()},
		Storages:    []*pb.Storage{NewSSD(), NewHDD()},
		Screen:      NewScreen(),
		Keyboard:    NewKeyboard(),
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   timestamppb.New(time.Now()),
	}
	return laptop
}
func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}