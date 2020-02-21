package main

// Server is a Test server
type Server struct {
	UnimplementedTestServer
}

// NewServer returns a new Server
func NewServer() *Server {
	s := &Server{}
	return s
}
