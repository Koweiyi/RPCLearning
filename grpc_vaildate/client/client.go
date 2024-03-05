package main

import (
	"RPCLearning/grpc_vaildate/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:50052", opts...)
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	client := proto.NewGreeterClient(conn)

	hello, err := client.SayHello(context.Background(), &proto.Person{
		Id:     1000,
		Email:  "koweiyi@gmail.com",
		Mobile: "17633611111",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(hello)
}
