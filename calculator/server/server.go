package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	calc "github.com/chazuka/hello-grpc/calculator"
)

const (
	address = "0.0.0.0:50051"
)

func main() {
	log.Println("connecting ...")
	connection, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	s := grpc.NewServer()
	calc.RegisterCalculatorServiceServer(s, &CalService{})
	log.Println("serving grpc service ...")
	if err := s.Serve(connection); err != nil {
		log.Fatal(err)
	}
}
