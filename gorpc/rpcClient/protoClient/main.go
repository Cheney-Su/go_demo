package main

import (
	"google.golang.org/grpc"
	"log"
	"go/gorpc/hello.pb"
	"golang.org/x/net/context"
	"fmt"
)

const (
	host = "127.0.0.1:1234"
)

func main() {

	client, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("grpc connect err...")
	}
	c := hello.NewGreeterClient(client)
	name := "shq"
	r, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: name})
	if err != nil {
		log.Fatalln("grpc sayHello err...")
	}
	fmt.Println("Greeting:", r.Message)
}