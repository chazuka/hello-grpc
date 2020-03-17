package main

import (
	"context"
	"errors"
	"io"
	"log"
	"sync"
	"time"

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

func (sdk *GreetingSDK) GreetClientServerStream(ctx context.Context, persons []pkg.Person, cbr func(string, error), cbs func(pkg.Person, error)) error {
	ss, err := sdk.connection.GreetClientServerStream(ctx)
	if err != nil {
		return err
	}

	wc := make(chan struct{})
	go func() {
		for i, p := range persons {
			log.Printf("sending message to server: %d. %s %s", i, p.GetFirstName(), p.GetLastName())
			if err := ss.Send(&pkg.GreetClientServerStreamRequest{Person: &p}); err != nil {
				cbs(p, err)
			}
			time.Sleep(1 * time.Second)
		}
		if err := ss.CloseSend(); err != nil {
			cbs(pkg.Person{}, err)
		}
	}()

	go func() {
		for {
			rs, err := ss.Recv()
			if errors.Is(err, io.EOF) {
				break
			}
			cbr(rs.GetGreeting(), err)
		}
		close(wc)
	}()
	<-wc
	return nil
}
