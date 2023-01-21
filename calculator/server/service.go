package main

import (
	"context"
	"errors"
	"io"
	"log"

	"github.com/chazuka/hello-grpc/calculator/pkg"
)

type CalService struct {
	pkg.CalculatorServiceServer
}

func (s *CalService) Addition(ctx context.Context, r *pkg.AdditionRequest) (*pkg.AdditionResponse, error) {
	res := r.GetFirst() + r.GetSecond()
	return &pkg.AdditionResponse{Result: res}, nil
}

func (s *CalService) PrimeNumberDecomposition(r *pkg.PrimeNumberDecompositionRequest, ss pkg.CalculatorService_PrimeNumberDecompositionServer) error {
	base := int32(2)
	number := r.GetNumber()
	for {
		if number < base {
			break
		}
		if number%base == 0 {
			_ = ss.Send(&pkg.PrimeNumberDecompositionResponse{Factor: base})
			number = number / base
			continue
		}
		base = base + 1
	}
	return nil
}

func (s *CalService) Average(ss pkg.CalculatorService_AverageServer) error {
	numbers := make([]int32, 0)
	for {
		r, err := ss.Recv()
		if errors.Is(err, io.EOF) {
			avg := func(n []int32) float64 {
				sum := int32(0)
				for _, i := range n {
					sum += i
				}
				return float64(sum) / float64(len(n))
			}(numbers)

			if err := ss.SendAndClose(&pkg.AverageResponse{Average: avg}); err != nil {
				return err
			}
			return nil
		}
		if err != nil {
			log.Printf("error receiving message from client")
			continue
		}
		numbers = append(numbers, r.GetNumber())
	}
}

func (s *CalService) FindMaximum(ss pkg.CalculatorService_FindMaximumServer) error {
	maximum := int32(0)
	for {
		r, err := ss.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			log.Println("error while receiving message from client")
			continue
		}
		if r.GetNumber() > maximum {
			maximum = r.GetNumber()
			if err := ss.Send(&pkg.FindMaximumResponse{Number: maximum}); err != nil {
				log.Println("error while sending message to client")
			}
		}
	}
}
