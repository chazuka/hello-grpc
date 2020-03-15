package main

import (
	"context"
	"errors"
	"io"
	"log"
	"sync"

	"github.com/chazuka/hello-grpc/greet/pkg"
	"google.golang.org/grpc"
)

//GreetingSDK structured client functionality of greeting service
type GreetingSDK struct {
	connection pkg.GreetServiceClient
}

func NewGreetingSDK(c grpc.ClientConnInterface) *GreetingSDK {
	return &GreetingSDK{
		connection: pkg.NewGreetServiceClient(c),
	}
}

func (sdk *GreetingSDK) Greet(ctx context.Context, person pkg.Person) (*pkg.GreetResponse, error) {
	result, err := sdk.connection.Greet(context.TODO(), &pkg.GreetRequest{Person: &person})
	return result, err
}

func (sdk *GreetingSDK) GreetServerStream(ctx context.Context, person pkg.Person, num int32, cb func(*pkg.GreetServerStreamResponse, error)) error {
	rs, err := sdk.connection.GreetServerStream(ctx, &pkg.GreetServerStreamRequest{Person: &person, Number: num})
	if err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	wg.Add(1)

	for {
		greeting, err := rs.Recv()
		if errors.Is(err, io.EOF) {
			wg.Done()
			break
		}
		cb(greeting, err)
	}
	wg.Wait()
	return nil
}

func (sdk *GreetingSDK) GreetClientStream(ctx context.Context, persons []pkg.Person) (*pkg.GreetClientStreamResponse, error) {
	ss, err := sdk.connection.GreetClientStream(ctx)
	if err != nil {
		return nil, err
	}
	for _, p := range persons {
		if err := ss.Send(&pkg.GreetClientStreamRequest{Person: &p}); err != nil {
			log.Printf("error sending: %s %s", p.GetFirstName(), p.GetLastName())
		}
	}
	return ss.CloseAndRecv()
}
