package ctrl

import (
	"context"
	"fmt"
	"grpc-starter/internal/service"
	userProto "grpc-starter/proto"
	"log"
)

type AuthController struct {
	userProto.UnimplementedAuthServiceServer
}

func NewAuthCtrl() *AuthController {
	return &AuthController{}
}

func validateLoginRequest(request *userProto.LoginRequest) error {
	return nil
}

func (a *AuthController) Login(ctx context.Context, request *userProto.LoginRequest) (*userProto.LoginResponse, error) {
	log.Println("user login -> ", request.Username, request.Password)
	// biz .
	user, err := service.NewAuthService().Login(ctx, request.Username, request.Password)
	if err != nil {
		// 错误日志
		return nil, err
	}
	// 做完了业务. 组装一些响应数据.

	resp := &userProto.LoginResponse{
		Token: "123",
		User: &userProto.User{
			Id: user.Id,
		},
	}
	// 最终再返回你的参数.
	return resp, nil
}

func (a *AuthController) Register(ctx context.Context, req *userProto.RegisterRequest) (*userProto.RegisterResponse, error) {
	userinfo, err := userFromContext(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(userinfo)
	return &userProto.RegisterResponse{}, nil
}
