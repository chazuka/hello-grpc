package main

import (
	"context"

	calc "github.com/chazuka/hello-grpc/calculator"
)

//CalService structured functionalities of calculator service
type CalService struct{}

//Addition add 2 numbers
func (s *CalService) Addition(ctx context.Context, r *calc.AdditionRequest) (*calc.AdditionResponse, error) {
	res := r.GetFirst() + r.GetSecond()
	return &calc.AdditionResponse{Result: res}, nil
}

func (s *CalService) PrimeNumberDecomposition(r *calc.PrimeNumberDecompositionRequest, ss calc.CalculatorService_PrimeNumberDecompositionServer) error {
	base := int32(2)
	number := r.GetNumber()
	for {
		if number < 1 {
			break
		}
		if number%base == 0 {
			ss.Send(&calc.PrimeNumberDecompositionResponse{Factor: base})
			number = number / base
			continue
		}
		base = base + 1

	}
	return nil
}
