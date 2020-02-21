package ossa

import (
	context "context"
	"log"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StreamIntercept intercepts stream GRPC traffic
func StreamIntercept(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := ss.Context()
	err := status.Errorf(codes.Unimplemented, "received stream intercept message with context: %v", ctx)
	log.Println(err)
	return err
}

// UnaryIntercept intercepts unary GRPC traffic
func UnaryIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	err = status.Errorf(codes.Aborted, "received stream intercept message with context: %v", ctx)
	log.Println(err)
	return nil, err
}
