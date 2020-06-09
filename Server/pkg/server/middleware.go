package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
)

// StreamLogger implements the StreamServerInterceptor interface. This function is used to log information like the IP
// address of the clients.
func StreamLogger(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	p, ok := peer.FromContext(ss.Context())
	if !ok {
		log.Println(fmt.Sprintf("Streaming RPC: %s was called by unknown.", info.FullMethod))
	}
	log.Println(fmt.Sprintf("IP: %s called Streaming RPC: %s using AuthType: %s", p.Addr.String(), info.FullMethod, p.AuthInfo.AuthType()))
	return handler(srv, ss)
}

// UnaryLogger implements the UnaryServerInterceptor interface. This function is used to log information like the IP
// address of the clients.
func UnaryLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		log.Println(fmt.Sprintf("RPC: %s was called by unknown.", info.FullMethod))
	}
	log.Println(fmt.Sprintf("IP: %s called RPC: %s using AuthType: %s", p.Addr.String(), info.FullMethod, p.AuthInfo.AuthType()))
	return handler(ctx, req)
}
