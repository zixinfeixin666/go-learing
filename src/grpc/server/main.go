package server

import (
	"context"
	"fmt"
	hello_grpc2 "go_test/src/grpc/pub"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello_grpc2.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hello_grpc2.Req) (res *hello_grpc2.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc2.Res{Message: "我是从服务端返回的grpc的内容"}, nil
}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	hello_grpc2.RegisterHelloGRPCServer(s, &server{})
	s.Serve(l)
}
