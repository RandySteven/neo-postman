package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal(err)
		return
	}
	opt := []grpc.ServerOption{
		grpc.MaxSendMsgSize(1024 * 1024 * 10),
		grpc.MaxRecvMsgSize(1024 * 1024 * 10),
		grpc.ChainStreamInterceptor(),
	}
	server := grpc.NewServer(opt...)

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
		return
	}
}
