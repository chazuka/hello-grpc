package main

import (
	"context"

	"github.com/chazuka/hello-grpc/calculator/pkg"
)

type CalService struct{}

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
