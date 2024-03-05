package main

import (
	"RPCLearning/new_helloworld/client_proxy"
	"fmt"
)

func main() {
	// 1.建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")

	var reply string
	err := client.Hello("koweiyi", &reply)
	if err != nil {
		panic("调用失败")
	}

	fmt.Println(reply)
}
