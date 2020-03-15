package main

import (
	"context"
	"log"

	"github.com/chazuka/hello-grpc/greet/pkg"
	"google.golang.org/grpc"
)

func unary(sdk *GreetingSDK, person pkg.Person) {
	result, err := sdk.Greet(context.TODO(), person)
	if err != nil {
		log.Printf("error while calling greet %v", err)
	}
	log.Printf("greeting from server: %s", result.GetGreeting())
}

func serverStream(sdk *GreetingSDK, person pkg.Person, num int) {
	cb := func(rw *pkg.GreetServerStreamResponse, err error) {
		if err != nil {
			log.Printf("error: %+v", err)
			return
		}
		log.Printf("greeting from server: %s", rw.GetGreeting())
	}
	err := sdk.GreetServerStream(context.TODO(), person, int32(num), cb)
	if err != nil {
		log.Fatal(err)
	}
}

func clientStream(sdk *GreetingSDK, persons []pkg.Person) {
	rw, err := sdk.GreetClientStream(context.TODO(), persons)
	if err != nil {
		log.Fatal(err)
	}

	greetings := rw.GetGreeting()
	log.Printf("received %d greetings from server", len(greetings))
	for i, g := range greetings {
		log.Printf("%d. %s", i, g)
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

	person := pkg.Person{
		FirstName: "Komang",
		LastName:  "Arthayasa",
	}
	unary(sdk, person)
	serverStream(sdk, person, 5)
	clientStream(sdk, []pkg.Person{
		{FirstName: "Komang", LastName: "Arthayasa"},
		{FirstName: "Altona", LastName: "Widjaja"},
		{FirstName: "Leo", LastName: "Rojito"},
		{FirstName: "Cokorda", LastName: "Susila"},
		{FirstName: "Steven", LastName: "Pangemanan"},
	})
}
