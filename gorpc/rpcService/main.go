package main

import (
	"net/rpc"
	"net"
	"log"
	"net/http"
)

type Rpc int

func (r *Rpc) Hi(text string, reply *string) error {
	*reply = "Hi," + text
	return nil
}

func main() {
	rpc.Register(new(Rpc))
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	log.Println("rpc servering...")
	http.Serve(l, nil)
}
