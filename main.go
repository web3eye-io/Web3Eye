package main

import (
	"context"

	"github.com/web3eye-io/Web3Eye/gen-car/pkg/car"
)

func main() {
	car.CreateCar(context.Background(), "/home/coast/cybertracer/go-car/cmd/car/list.car", []string{"/home/coast/cybertracer/go-car/cmd/car/list.go"}, car.DefaultCarVersion)
}
