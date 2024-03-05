package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(req string, resp *string) error {
	*resp = "hello, " + req
	return nil
}

func main() {

	//1. 实例化一个server
	lis, _ := net.Listen("tcp", ":1234")

	// 2.注册处理逻辑handler
	_ = rpc.RegisterName("HelloService", &HelloService{})

	// 3.开启服务
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic("连接失败")
		}
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
