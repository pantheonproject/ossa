package main

import (
	"net"

	"google.golang.org/grpc"
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
