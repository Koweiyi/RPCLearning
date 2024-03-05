package main

import (
	"RPCLearning/grpc_vaildate/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(_ context.Context, req *proto.Person) (*proto.Person, error) {
	return &proto.Person{
		Id:    req.Id + 1,
		Email: "hello " + req.Email,
	}, nil
}

type Validator interface {
	Validate() error
}

func main() {
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if r, ok := req.(Validator); ok {
			if err := r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	}
	var opts []grpc.ServerOption

	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	g := grpc.NewServer(opts...)

	proto.RegisterGreeterServer(g, &Server{})

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}

	err = g.Serve(lis)
	if err != nil {
		panic(err)
	}
}
