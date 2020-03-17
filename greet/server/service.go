package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/chazuka/hello-grpc/greet/pkg"
)

type GreetService struct{}

func (s *GreetService) Greet(ctx context.Context, r *pkg.GreetRequest) (*pkg.GreetResponse, error) {
	log.Printf("serving Greet request ...")
	fn := r.GetPerson().GetFirstName()
	ln := r.GetPerson().GetLastName()
	rw := &pkg.GreetResponse{
		Greeting: fmt.Sprintf("Hello %s %s", fn, ln),
	}

	return rw, nil
}

func (s *GreetService) GreetServerStream(r *pkg.GreetServerStreamRequest, ss pkg.GreetService_GreetServerStreamServer) error {
	fn := r.GetPerson().GetFirstName()
	ln := r.GetPerson().GetLastName()
	num := int(r.GetNumber())
	if num < 1 {
		num = 1
	}
	log.Printf("stream greeting %s %s %d times", fn, ln, num)

	for i := 1; i <= num; i++ {
		rw := &pkg.GreetServerStreamResponse{
			Greeting: fmt.Sprintf("Hello %s %s #%d", fn, ln, i),
		}
		log.Printf("sending response %d", i)
		if err := ss.Send(rw); err != nil {
			return err
		}
	}

	return nil
}

func (s *GreetService) GreetClientStream(ss pkg.GreetService_GreetClientStreamServer) error {
	log.Printf("serving GreetClientStream ...")
	greetings := make([]string, 0)
	for {
		log.Printf("receiving message from client")
		r, err := ss.Recv()
		if errors.Is(err, io.EOF) {
			log.Printf("all message received, sending response")
			return ss.SendAndClose(&pkg.GreetClientStreamResponse{Greeting: greetings})
		}
		if err != nil {
			return err
		}
		greetings = append(greetings, fmt.Sprintf("Hello %s %s", r.GetPerson().GetFirstName(), r.GetPerson().GetLastName()))
	}
}

func (s *GreetService) GreetClientServerStream(ss pkg.GreetService_GreetClientServerStreamServer) error {
	log.Printf("serving GreetClientServerStream ...")
	for {
		r, err := ss.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		p := r.GetPerson()
		if err = ss.Send(&pkg.GreetClientServerStreamResponse{Greeting: fmt.Sprintf("Hello %s %s", p.GetFirstName(), p.GetLastName())}); err != nil {
			log.Printf("error while sending message")
		}
	}
}
