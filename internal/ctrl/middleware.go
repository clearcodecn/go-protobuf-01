package ctrl

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-starter/internal/model"
)

func (AuthController) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if fullMethodName == "/proto.AuthService/Login" {
		return ctx, nil
	}

	// 1. 从ctx 取出 token信息，
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Error(codes.Unauthenticated, "noAuthorized")
	}
	token := md.Get("token")

	key := md.Get("key")
	fmt.Println("key -> ", key)
	if len(token) > 0 {
		tk := token[0]
		// 2. 验证token
		if tk == "123" {
			// 3. 拦截

			ctx = newUserContext(ctx, &model.UserModel{
				Id:       1,
				Username: "admin",
				Password: "admin",
			})
			return ctx, nil
		}
	}
	// 放过
	return ctx, status.Error(codes.Unauthenticated, "noAuthorized")
}

type userInfoKey struct{}

func newUserContext(ctx context.Context, userinfo *model.UserModel) context.Context {
	return context.WithValue(ctx, userInfoKey{}, userinfo)
}

func userFromContext(ctx context.Context) (*model.UserModel, error) {
	um, ok := ctx.Value(userInfoKey{}).(*model.UserModel)
	if !ok {
		return nil, errors.New("not found")
	}
	return um, nil
}
