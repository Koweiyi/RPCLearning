package main

import (
	"RPCLearning/grpc_stream/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

const PORT = ":50052"

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	for i := 0; i < 10; i++ {
		err := res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (s *server) PutStream(streamServer proto.Greeter_PutStreamServer) error {
	for {
		if a, err := streamServer.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}

func (s *server) AllStream(streamServer proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := streamServer.Recv()
			fmt.Println("收到客户端消息: " + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			var text string
			_, _ = fmt.Scanln(&text)
			_ = streamServer.Send(&proto.StreamResData{
				Data: text,
			})

		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}

	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &server{})
	err = g.Serve(lis)
	if err != nil {
		panic(err)
	}
}
