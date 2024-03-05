package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, _ := rpc.Dial("tcp", "localhost:1234")

	var res string
	err := client.Call("HelloService.Hello", "koweiyi", &res)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(res)
}
