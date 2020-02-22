package ossa

import (
	context "context"
	fmt "fmt"
	"runtime/debug"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Auditor intercepts messages, processes audit logging rules, exports relevant events and messages and configured, then proceeds
type Auditor struct {
	logger Logger
}

// NewAuditor creates an Auditor
func NewAuditor(logger Logger) *Auditor {
	a := &Auditor{
		logger: logger,
	}
	return a
}

// StreamServerIntercept intercepts inbound stream GRPC traffic
func (a *Auditor) StreamServerIntercept(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	var txid string
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(codes.Internal, "ossa: caught panic in StreamServerIntercept: %v", r)
			go a.getLogger().LogErrorf(txid, "ossa: caught panic in StreamServerIntercept: %v, %s", r, debug.Stack())
		}
	}()

	// Not implemented, pass through
	return handler(srv, ss)
}

// UnaryServerIntercept intercepts inbound unary GRPC traffic
func (a *Auditor) UnaryServerIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	var txid string
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(codes.Internal, "ossa: caught panic in UnaryServerIntercept: %v", r)
			go a.getLogger().LogErrorf(txid, "ossa: caught panic in UnaryServerIntercept: %v, %s", r, debug.Stack())
		}
	}()
	preq, ok := req.(proto.Message)
	// fmt.Println(ok)
	exts, err := proto.GetExtension(preq, E_Sensitive)
	if err != nil {
		fmt.Println(exts)
	}
	if ok {
		mType := proto.MessageType(proto.MessageName(preq)).Elem()
		fmt.Println(mType)
		for i := 0; i < mType.NumField(); i++ {
			fmt.Println(mType.Field(i))
		}
	}
	// fmt.Println(info)
	// a.getLogger().LogDebugf(txid, "req: %+v", req)
	// Not implemented, pass through
	return handler(ctx, req)
}

// UnaryClientIntercept intercepts outbound unary GRPC traffic
func (a *Auditor) UnaryClientIntercept(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	var txid string
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(codes.Internal, "ossa: caught panic in UnaryClientIntercept: %v", r)
			go a.getLogger().LogErrorf(txid, "ossa: caught panic in UnaryClientIntercept: %v, %s", r, debug.Stack())
		}
	}()
	// Not implemented, pass through
	return invoker(ctx, method, req, reply, cc, opts...)
}

// StreamClientInterceptor intercepts outbound stream GRPC traffic
func (a *Auditor) StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (cs grpc.ClientStream, err error) {
	var txid string
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(codes.Internal, "ossa: caught panic in StreamClientInterceptor: %v", r)
			go a.getLogger().LogErrorf(txid, "ossa: caught panic in StreamClientInterceptor: %v, %s", r, debug.Stack())
		}
	}()
	// Not implemented, pass through
	return streamer(ctx, desc, cc, method, opts...)
}

func (a *Auditor) getLogger() Logger {
	if a == nil || a.logger == nil {
		return &noLogger{}
	}
	return a.logger
}
