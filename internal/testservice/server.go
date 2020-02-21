package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server is a Test server
type Server struct {
	UnimplementedTestServer
	server *grpc.Server
}

// NewServer returns a new Server
func NewServer(server *grpc.Server) *Server {
	s := &Server{
		server: server,
	}
	RegisterTestServer(server, s)
	return s
}

// Serve starts the server
func (s *Server) Serve(lis net.Listener) error {
	return s.server.Serve(lis)
}

// AddThing adds a thing
func (s *Server) AddThing(ctx context.Context, req *AddThingRequest) (*AddThingResponse, error) {
	return &AddThingResponse{}, status.Error(codes.Canceled, "test")
}
