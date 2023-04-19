package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-starter/internal/ctrl"
	userProto "grpc-starter/proto"
	"net"
)

func StartGRPCServer() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	authCtrl := ctrl.NewAuthCtrl()
	userProto.RegisterAuthServiceServer(s, authCtrl)

	err = s.Serve(ln)
	if err != nil {
		fmt.Println(err)
		return
	}
}
