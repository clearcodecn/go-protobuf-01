package grpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
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

	s := grpc.NewServer(
		//grpc.ChainUnaryInterceptor(
		//	func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//		log.Println("login before")
		//		resp, err = handler(ctx, req)
		//		log.Println("login after")
		//		return resp, err
		//	},
		//
		//	func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//		log.Println("login2 before")
		//		resp, err = handler(ctx, req)
		//		log.Println("login2 after")
		//		return resp, err
		//	},
		//
		//	func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//		log.Println("login3 before")
		//		resp, err = handler(ctx, req)
		//		log.Println("login3 after")
		//		return resp, err
		//	},
		//),

		// 期望： key:value  key = token ,  value = jwt / 123
		grpc.ChainUnaryInterceptor(
			auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
				return ctx, nil
			}),
			validator.UnaryServerInterceptor(validator.WithFailFast()),
		),
	)

	authCtrl := ctrl.NewAuthCtrl()
	userProto.RegisterAuthServiceServer(s, authCtrl)

	err = s.Serve(ln)
	if err != nil {
		fmt.Println(err)
		return
	}
}
