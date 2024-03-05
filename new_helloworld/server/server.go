package main

import (
	"RPCLearning/new_helloworld/handler"
	"RPCLearning/new_helloworld/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	//1. 实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	// 2.注册处理逻辑handler
	_ = server_proxy.RegisterHelloService(&handler.HelloService{})
	// 3.开启服务
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
