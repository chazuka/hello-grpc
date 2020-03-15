package main

import (
	"context"
	"log"

	"github.com/chazuka/hello-grpc/greet"

	"google.golang.org/grpc"
)

func once(sdk *GreetingSDK, person greet.Person) {
	result, err := sdk.Greet(context.TODO(), person)
	if err != nil {
		log.Printf("error while calling greet %v", err)
	}
	log.Printf("greeting from server: %s", result.GetGreeting())
}

func multiple(sdk *GreetingSDK, person greet.Person) {
	cb := func(rw *greet.GreetStreamResponse, err error) {
		if err != nil {
			log.Printf("error: %+v", err)
			return
		}
		log.Printf("greeting from server: %s", rw.GetGreeting())
	}
	err := sdk.GreetStream(context.TODO(), person, cb)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("connecting ....")
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	sdk := NewGreetingSDK(connection)

	person := greet.Person{
		FirstName: "Komang",
		LastName:  "Arthayasa",
	}
	once(sdk, person)
	multiple(sdk, person)
}
