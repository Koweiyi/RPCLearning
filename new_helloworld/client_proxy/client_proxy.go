package client_proxy

import (
	"RPCLearning/new_helloworld/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(network, address string) HelloServiceStub {
	client, err := rpc.Dial(network, address)
	if err != nil {
		panic("connect error!")
	}
	return HelloServiceStub{client}
}

func (c *HelloServiceStub) Hello(req string, rep *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", req, rep)
	if err != nil {
		return err
	}
	return nil
}
