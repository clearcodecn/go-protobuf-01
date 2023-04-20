package proto

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *LoginRequest) Validate() error {
	if r.Username == "" {
		return status.Error(codes.Aborted, "用户名不能空")
	}
	return nil
}
