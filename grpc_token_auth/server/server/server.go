package main

import (
	"RPCLearning/grpc_token_auth/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "hello " + req.GetName(),
	}, nil
}

func main() {

	// 设置拦截器
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("接收到新的请求")

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}

		//fmt.Println(md)
		var (
			appid  string
			appkey string
		)
		if appidSlice, ok := md["appid"]; ok {
			appid = appidSlice[0]
		}
		if appkeySlice, ok := md["appkey"]; ok {
			appkey = appkeySlice[0]
		}
		//fmt.Println(appid, appkey)

		if appid != "101010" || appkey != "a secret" {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}

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
