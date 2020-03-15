package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	calc "github.com/chazuka/hello-grpc/calculator"
)

//CalculatorServer structured functionalities of calculator service
type CalculatorServer struct{}

//Addition add 2 numbers
func (s *CalculatorServer) Addition(ctx context.Context, r *calc.AdditionRequest) (*calc.AdditionResponse, error) {
	res := r.GetFirst() + r.GetSecond()
	return &calc.AdditionResponse{Result: res}, nil
}

func main() {
	// listen to TCP connection
	log.Println("connecting to TCP ...")
	connection, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}
	// build & configure grpc server
	s := grpc.NewServer()
	// attach handler
	calc.RegisterCalculatorServiceServer(s, &CalculatorServer{})
	// serve grpc service
	log.Println("serving grpc service ...")
	if err := s.Serve(connection); err != nil {
		log.Fatal(err)
	}
}
