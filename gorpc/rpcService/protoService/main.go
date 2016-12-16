package main

import (
	"golang.org/x/net/context"
	"go/gorpc/hello.pb"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	host = ":1234"
)

type server struct {

}

func (s *server) SayHello(c context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message:"Hello " + in.Name}, nil
}

func main() {

	listen, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalln("listen error...")
	}
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	log.Println("grpc servicing...")
	if err = s.Serve(listen); err != nil {
		log.Fatalln("failed to serve:", err)
	}

}