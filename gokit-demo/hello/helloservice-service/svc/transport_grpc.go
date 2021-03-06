// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 574fb16d86
// Version Date: 2019年 4月12日 星期五 00时42分59秒 UTC

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/akka/gokit-demo/hello"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC HelloServiceServer.
func MakeGRPCServer(endpoints Endpoints) pb.HelloServiceServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	return &grpcServer{
		// helloservice

		sayhello: grpctransport.NewServer(
			endpoints.SayHelloEndpoint,
			DecodeGRPCSayHelloRequest,
			EncodeGRPCSayHelloResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the HelloServiceServer interface
type grpcServer struct {
	sayhello grpctransport.Handler
}

// Methods for grpcServer to implement HelloServiceServer interface

func (s *grpcServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	_, rep, err := s.sayhello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HelloResponse), nil
}

// Server Decode

// DecodeGRPCSayHelloRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC sayhello request to a user-domain sayhello request. Primarily useful in a server.
func DecodeGRPCSayHelloRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.HelloRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCSayHelloResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain sayhello response to a gRPC sayhello reply. Primarily useful in a server.
func EncodeGRPCSayHelloResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.HelloResponse)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
