package main

import (
	"context"
	"errors"
	"io"
	"sync"

	"github.com/chazuka/hello-grpc/greet/pkg"
	"google.golang.org/grpc"
)

//GreetingSDK structured client functionality of greeting service
type GreetingSDK struct {
	connection pkg.GreetServiceClient
}

//NewGreetingSDK it build a new instance of GreetingSDK
func NewGreetingSDK(c grpc.ClientConnInterface) *GreetingSDK {
	return &GreetingSDK{
		connection: pkg.NewGreetServiceClient(c),
	}
}

//Greet request single greeting for a given person
func (sdk *GreetingSDK) Greet(ctx context.Context, person pkg.Person) (*pkg.GreetResponse, error) {
	result, err := sdk.connection.Greet(context.TODO(), &pkg.GreetRequest{Person: &person})
	return result, err
}

//GreetStream request stream greeting for a person
func (sdk *GreetingSDK) GreetStream(ctx context.Context, person pkg.Person, cb func(*pkg.GreetStreamResponse, error)) error {
	rs, err := sdk.connection.GreetStream(ctx, &pkg.GreetStreamRequest{Person: &person})
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
