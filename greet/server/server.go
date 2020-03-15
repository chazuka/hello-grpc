package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/chazuka/hello-grpc/greet"
	"google.golang.org/grpc"
)

type greetService struct{}

func (s *greetService) Greet(ctx context.Context, r *greet.GreetRequest) (*greet.GreetResponse, error) {
	fn := r.GetGreeting().GetFirstName()
	ln := r.GetGreeting().GetLastName()
	log.Printf("serving greeting request with payload %s and %s", fn, ln)
	rw := &greet.GreetResponse{
		Greeting: fmt.Sprintf("Hello %s %s", fn, ln),
	}

	return rw, nil
}

func main() {
	log.Println("listening to TCP connection")
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	greetServer := grpc.NewServer()
	greet.RegisterGreetServiceServer(greetServer, &greetService{})
	log.Println("serving request")
	if err := greetServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
