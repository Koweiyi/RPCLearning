package main

import (
	"RPCLearning/grpc_token_auth/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "a secret",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {

	// 设置拦截器
	//interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	start := time.Now()
	//	md := metadata.New(map[string]string{
	//		"appid":  "101010",
	//		"appkey": "a secret",
	//	})
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	fmt.Printf("耗时：%s\n", time.Since(start))
	//	return err
	//}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	opts = append(opts, grpc.WithPerRPCCredentials(customCredential{}))

	// 建立连接并通过连接实例化一个客户端对象
	conn, err := grpc.Dial("localhost:50052", opts...)
	if err != nil {
		panic(err)
	}
	client := proto.NewGreeterClient(conn)

	// 调用服务
	helloReply, err := client.SayHello(context.Background(), &proto.HelloRequest{
		Name: "koweiyi",
	})

	// 解析错误
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic("解析错误失败！")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
	fmt.Println(helloReply.Message)
}
