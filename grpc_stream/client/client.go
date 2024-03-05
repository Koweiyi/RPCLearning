package main

import (
	"RPCLearning/grpc_stream/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
)

func main() {

	//// 服务端流模式
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	c := proto.NewGreeterClient(conn)
	//res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "koweiyi"})
	//
	//for {
	//	recv, err := res.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(recv.GetData())
	//}
	//
	//// 客户端流模式
	//
	//putS, _ := c.PutStream(context.Background())
	//for i := 0; i < 10; i++ {
	//	err := putS.Send(&proto.StreamReqData{
	//		Data: fmt.Sprintf("koweiyi%d", i),
	//	})
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	time.Sleep(time.Second)
	//}

	streamServer, _ := c.AllStream(context.Background())
	// 双向流模式
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := streamServer.Recv()
			fmt.Println("收到服务端消息: " + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			var text string
			_, _ = fmt.Scanln(&text)
			_ = streamServer.Send(&proto.StreamReqData{
				Data: text,
			})

		}
	}()
	wg.Wait()

}
