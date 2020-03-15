package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chazuka/hello-grpc/greet/pkg"
)

type GreetService struct{}

func (s *GreetService) Greet(ctx context.Context, r *pkg.GreetRequest) (*pkg.GreetResponse, error) {
	fn := r.GetPerson().GetFirstName()
	ln := r.GetPerson().GetLastName()
	log.Printf("greeting a person with payload %s and %s", fn, ln)
	rw := &pkg.GreetResponse{
		Greeting: fmt.Sprintf("Hello %s %s", fn, ln),
	}

	return rw, nil
}

func (s *GreetService) GreetStream(r *pkg.GreetStreamRequest, ss pkg.GreetService_GreetStreamServer) error {
	fn := r.GetPerson().GetFirstName()
	ln := r.GetPerson().GetLastName()
	log.Printf("stream greeting a person with payload %s and %s", fn, ln)

	for i := 0; i < 10; i++ {
		rw := &pkg.GreetStreamResponse{
			Greeting: fmt.Sprintf("Hello %s %s #%d", fn, ln, i),
		}
		log.Printf("sending response %d", i)
		if err := ss.Send(rw); err != nil {
			return err
		}
		// emulate delay process
		time.Sleep(1 * time.Second)
	}

	return nil
}
