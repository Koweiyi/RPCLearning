package main

import (
	"RPCLearning/grpc_test/proto"
	"context"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	time.Sleep(time.Second * 4)
	return &proto.HelloReply{
		Message: "hello " + req.Name,
	}, nil
}

func main() {
	// 实例化一个server
	g := grpc.NewServer()

	// 注册服务
	proto.RegisterGreeterServer(g, &Server{})

	// 监听端口

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	err = g.Serve(lis)
	if err != nil {
		panic(err)
	}
}
