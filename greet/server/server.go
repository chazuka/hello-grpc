package main

import (
	"log"
	"net"

	"github.com/chazuka/hello-grpc/greet/pkg"
	"google.golang.org/grpc"
)

const (
	address = ":50051"
)

func main() {
	log.Println("opening TCP connection ...")

	// create new TCP listener
	connection, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	// create new GRPC server
	server := grpc.NewServer()

	// register greet service into grpc server
	pkg.RegisterGreetServiceServer(server, &GreetService{})

	log.Println("running server at ", address)

	// running grpc server
	if err := server.Serve(connection); err != nil {
		log.Fatal(err)
	}
}
