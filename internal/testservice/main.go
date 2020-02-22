package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/pantheonproject/ossa"
	"google.golang.org/grpc"
)

var bindAddress = "localhost:6687"

func main() {
	// Create the server
	auditor := ossa.NewAuditor(&ossa.StdOutLogger{})
	gserver := grpc.NewServer(grpc.UnaryInterceptor(auditor.UnaryServerIntercept), grpc.StreamInterceptor(auditor.StreamServerIntercept))
	s := NewServer(gserver)
	lis, err := net.Listen("tcp", bindAddress)
	if err != nil {
		log.Fatalf("error starting net listener: %v", err)
	}
	go serverStart(s, lis)
	time.Sleep(2 * time.Second)
	clientConnect(bindAddress)
}

func clientConnect(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error dialing client connection: %v", err)
	}
	client := NewTestClient(conn)
	ctx := context.Background()
	res, err := client.AddThing(ctx, &AddThingRequest{})
	if err != nil {
		log.Fatalf("error adding thing: %v", err)
	}
	log.Printf("response: %v\n", res)
}

func serverStart(s *Server, lis net.Listener) {
	err := s.Serve(lis)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
	log.Println("exited server with no error")
}
