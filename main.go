package main

import (

)

type Company struct{
	Name string
	Address string
}

type Employee struct {
	Name string
	Company Company
}

func RpcPrintln(e Employee) {
	/*
		客户端
			1.建立连接 tcp/http
			2.序列化 employee 对象
			3.发送json字符串
			4.等待服务器发送结果
			5.将服务器返回的数据解析成PrintResult对象 -反序列化
		服务端
			1.监听网络端口
			2.读数据 - 二进制的json数据
			3.将数据反序列化为 Employee 对象
			4.处理业务逻辑
			5.将结果序列化为json数据
			6.返回数据

		序列化和反序列化的协议是可选择的，不一定选择json
			也可能使用xml, protobuf, msgpack 等
		
	*/
}

func main() {
	
}