package main

import (
	"context"
	"log"

	"github.com/chazuka/hello-grpc/greet"

	"google.golang.org/grpc"
)

func main() {
	log.Println("I am a grpc client")
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	greetClient := greet.NewGreetServiceClient(connection)

	result, err := greetClient.Greet(context.TODO(), &greet.GreetRequest{Greeting: &greet.Greeting{
		FirstName: "Komang",
		LastName:  "Arthayasa",
	}})
	if err != nil {
		log.Printf("error while calling greet %v", err)
	}
	log.Printf("greeting from server: %s", result.GetGreeting())
}
