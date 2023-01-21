package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"google.golang.org/grpc"
)

const (
	address = ":50051"
)

func main() {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = connection.Close()
	}()

	sdk := NewCalculator(connection)
	handleAddition(sdk, 3, 10)
	handlePrimeNumberDecomposition(sdk, 120)
	handleAverage(sdk, []int32{1, 2, 3, 4, 5, 6, 7})
	handleFindMaximum(sdk, []int32{2, 1, 3, 6, 6, 5, 8})

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

func handleAverage(sdk *CalculatorSDK, numbers []int32) {
	cb := func(i int32, err error) {
		if err != nil {
			log.Printf("sending: %d - error: %+v", i, err)
			return
		}
		log.Printf("sending: %d - OK", i)
	}

	rw, err := sdk.Average(context.TODO(), numbers, cb)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("average of %+v is %.2f", numbers, rw)
}

func handleFindMaximum(sdk *CalculatorSDK, numbers []int32) {
	cbs := func(n int32, err error) {
		if err != nil {
			log.Printf("error sending: %+v", err)
			return
		}
		log.Printf("sent number: %d", n)
	}
	cbr := func(n int32, err error) {
		if err != nil {
			log.Printf("error receiving: %+v", err)
			return
		}
		log.Printf("receive number: %d", n)
	}
	if err := sdk.FindMaximum(context.Background(), numbers, cbr, cbs); err != nil {
		log.Println(err)
	}
}
