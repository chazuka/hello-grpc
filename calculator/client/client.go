package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/chazuka/hello-grpc/calculator"
)

func main() {
	// build grpc connection client
	connection, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	calc := NewCalculator(connection)
	res, _ := calc.Add(context.TODO(), 3, 10)
	log.Printf("result of %d + %d = %d", 3, 10, res)

}

type CalculatorSDK struct {
	connection calculator.CalculatorServiceClient
}

func NewCalculator(l grpc.ClientConnInterface) *CalculatorSDK {
	c := calculator.NewCalculatorServiceClient(l)
	return &CalculatorSDK{connection: c}
}

func (c *CalculatorSDK) Add(ctx context.Context, a, b int32) (int32, error) {
	r, err := c.connection.Addition(ctx, &calculator.AdditionRequest{First: a, Second: b})
	if err != nil {
		return 0, err
	}
	return r.Result, nil
}
