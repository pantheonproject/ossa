package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Create the server
	gserver := grpc.NewServer()
	s := NewServer(gserver)
	lis, err := net.Listen("tcp", "localhost:6687")
	if err != nil {
		log.Fatalf("error starting net listener: %v", err)
	}
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
