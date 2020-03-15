package main

import (
	"log"
	"net"

	"github.com/chazuka/hello-grpc/greet/pkg"
	"google.golang.org/grpc"
)

const (
	address = "0.0.0.0:50051"
)

func main() {
	log.Println("opening TCP connection ...")
	connection, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pkg.RegisterGreetServiceServer(server, &GreetService{})
	log.Println("running server ...")
	if err := server.Serve(connection); err != nil {
		log.Fatal(err)
	}
}
