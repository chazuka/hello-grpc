package main

import (
	"log"
	"net"

	"github.com/chazuka/hello-grpc/calculator/pkg"
	"google.golang.org/grpc"
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
	pkg.RegisterCalculatorServiceServer(s, &CalService{})
	log.Println("serving grpc service ...")
	if err := s.Serve(connection); err != nil {
		log.Fatal(err)
	}
}
