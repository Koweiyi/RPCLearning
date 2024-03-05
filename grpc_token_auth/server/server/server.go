package main

import (
	"RPCLearning/grpc_intepretor/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "hello " + req.GetName(),
	}, nil
}

func main() {

	// 设置拦截器
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("接收到新的请求")
		res, err := handler(ctx, req)
		fmt.Println("请求已完成")
		return res, err
	}
	opt := grpc.UnaryInterceptor(interceptor)

	// 1.实例化一个 server
	g := grpc.NewServer(opt)

	// 2.注册服务
	proto.RegisterGreeterServer(g, &server{})

	// 3.监听服务
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}

	// 4.开启服务
	err = g.Serve(lis)
	if err != nil {
		fmt.Println(err.Error())
	}
}
