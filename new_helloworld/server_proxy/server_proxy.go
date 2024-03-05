package server_proxy

import (
	"RPCLearning/new_helloworld/handler"
	"net/rpc"
)

type HelloServicer interface {
	Hello(req string, rep *string) error
}

func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
