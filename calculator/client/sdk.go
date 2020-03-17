package main

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/chazuka/hello-grpc/calculator/pkg"
	"google.golang.org/grpc"
)

type CalculatorSDK struct {
	connection pkg.CalculatorServiceClient
}

func NewCalculator(l grpc.ClientConnInterface) *CalculatorSDK {
	c := pkg.NewCalculatorServiceClient(l)
	return &CalculatorSDK{connection: c}
}

func (c *CalculatorSDK) Add(ctx context.Context, a, b int32) (int32, error) {
	r, err := c.connection.Addition(ctx, &pkg.AdditionRequest{First: a, Second: b})
	return r.GetResult(), err
}

//PrimeNumberDecomposition it handles stream response from server and return error accordingly
//it will send response process through callback function `cb`
func (c *CalculatorSDK) PrimeNumberDecomposition(ctx context.Context, n int32, cb func(int32, error)) error {
	ss, err := c.connection.PrimeNumberDecomposition(ctx, &pkg.PrimeNumberDecompositionRequest{Number: n})
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

//Average it handles streams number, return average number and error accordingly
//it notify each sending process through callback function `cb`
func (c *CalculatorSDK) Average(ctx context.Context, numbers []int32, cb func(int32, error)) (float64, error) {
	ss, err := c.connection.Average(ctx)
	if err != nil {
		return 0.0, err
	}
	for _, i := range numbers {
		err := ss.Send(&pkg.AverageRequest{Number: i})
		if cb != nil {
			cb(i, err)
		}
	}
	rw, err := ss.CloseAndRecv()
	if err != nil {
		return 0.0, err
	}
	return rw.GetAverage(), nil
}

func (c *CalculatorSDK) FindMaximum(ctx context.Context, numbers []int32, cbr func(int32, error), cbs func(int32, error)) error {
	ss, err := c.connection.FindMaximum(ctx)
	if err != nil {
		return err
	}

	wc := make(chan struct{})
	// sending
	go func() {
		for _, n := range numbers {
			if err := ss.Send(&pkg.FindMaximumRequest{Number: n}); err != nil {
				cbs(n, err)
				continue
			}
			cbs(n, nil)
			time.Sleep(1 * time.Second) // simulate delay
		}
		if err := ss.CloseSend(); err != nil {
			cbs(0, err)
		}
	}()

	// receiving
	go func() {
		for {
			r, err := ss.Recv()
			if errors.Is(err, io.EOF) {
				break
			}
			cbr(r.GetNumber(), err)
		}
		close(wc)
	}()

	<-wc

	return nil
}
