package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/NaveedZahoorKhan/golang-grpc-demo/pb"
	"github.com/jinzhu/copier"
)

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

var ErrAlreadyExist = errors.New("record already exist")

func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]

	if laptop == nil {
		return nil, nil
	}

	other := &pb.Laptop{}

	err := copier.Copy(other, laptop)

	if err != nil {
		return nil, err
	}

	return other, nil
}
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if store.data[laptop.Id] != nil {
		return ErrAlreadyExist
	}
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}
	store.data[other.Id] = other
	return nil
}
