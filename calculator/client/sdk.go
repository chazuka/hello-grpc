package main

import (
	"context"
	"errors"
	"io"
	"sync"

	"github.com/chazuka/hello-grpc/calculator"
	"google.golang.org/grpc"
)

type CalculatorSDK struct {
	connection calculator.CalculatorServiceClient
}

func NewCalculator(l grpc.ClientConnInterface) *CalculatorSDK {
	c := calculator.NewCalculatorServiceClient(l)
	return &CalculatorSDK{connection: c}
}

func (c *CalculatorSDK) Add(ctx context.Context, a, b int32) (int32, error) {
	r, err := c.connection.Addition(ctx, &calculator.AdditionRequest{First: a, Second: b})
	return r.GetResult(), err
}

func (c *CalculatorSDK) PrimeNumberDecomposition(ctx context.Context, n int32, cb func(int32, error)) error {
	ss, err := c.connection.PrimeNumberDecomposition(ctx, &calculator.PrimeNumberDecompositionRequest{Number: n})
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	for {
		r, err := ss.Recv()
		if errors.Is(err, io.EOF) {
			wg.Done()
			break
		}
		cb(r.GetFactor(), err)
	}
	wg.Wait()
	return nil
}
