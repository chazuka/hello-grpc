package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

const (
	address = "0.0.0.0:50051"
)

func main() {
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	sdk := NewCalculator(connection)
	handleAddition(sdk, 3, 10)
	handlePrimeNumberDecomposition(sdk, 120)

}

func handleAddition(sdk *CalculatorSDK, a, b int32) {
	res, _ := sdk.Add(context.TODO(), a, b)
	log.Printf("result of %d + %d = %d", a, b, res)
}

func handlePrimeNumberDecomposition(sdk *CalculatorSDK, n int32) {
	factors := make([]int32, 0)
	cb := func(f int32, err error) {
		if err != nil {
			log.Printf("error: %+v", err)
			return
		}
		factors = append(factors, f)
		log.Printf("response from server - %d is a factor", f)
	}
	if err := sdk.PrimeNumberDecomposition(context.TODO(), 120, cb); err != nil {
		log.Printf("error: %+v", err)
		return
	}
	log.Printf("factors of %d are: %v", n, factors)
}
