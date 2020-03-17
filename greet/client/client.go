package main

import (
	"context"
	"log"

	"github.com/chazuka/hello-grpc/greet/pkg"
	"google.golang.org/grpc"
)

func unary(sdk *GreetingSDK, person pkg.Person) {
	log.Println("requesting unary process")
	result, err := sdk.Greet(context.TODO(), person)
	if err != nil {
		log.Printf("error while calling greet %v", err)
	}
	log.Printf("greeting from server: %s", result.GetGreeting())
}

func serverStream(sdk *GreetingSDK, person pkg.Person, num int) {
	log.Println("requesting server stream process")
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
	log.Println("requesting client stream process")
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

func bidirectionalStream(sdk *GreetingSDK, persons []pkg.Person) {
	log.Println("requesting bidirectional stream process")
	cbr := func(r string, err error) {
		if err != nil {
			log.Printf("error receiving message from upstream: %+v", err)
			return
		}
		log.Printf("received: %s", r)
	}
	cbs := func(p pkg.Person, err error) {
		if err != nil {
			log.Printf("error sending message to upstream: %+v", err)
			return
		}
		log.Printf("sent: %s %s", p.GetFirstName(), p.GetLastName())
	}
	if err := sdk.GreetClientServerStream(context.Background(), persons, cbr, cbs); err != nil {
		log.Printf("error: %+v", err)
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
		{FirstName: "Leo", LastName: "Rodjito"},
		{FirstName: "Cokorda", LastName: "Susila"},
		{FirstName: "Steven", LastName: "Pangemanan"},
	})
	bidirectionalStream(sdk, []pkg.Person{
		{FirstName: "Komang", LastName: "Arthayasa"},
		{FirstName: "Altona", LastName: "Widjaja"},
		{FirstName: "Leo", LastName: "Rodjito"},
		{FirstName: "Cokorda", LastName: "Susila"},
		{FirstName: "Steven", LastName: "Pangemanan"},
	})
}
