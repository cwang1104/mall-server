package rpcUser

import (
	"context"
	pbUser "mall_user/proto/user"
)

type User struct {
	pbUser.UserHandler
}

//用户注册
func (u *User) UserRegister(ctx context.Context, in *pbUser.UserRequest, out *pbUser.UserResponse) error {

	return nil
}

func (u *User) UserSendEmail(context.Context, *pbUser.UserMailRequest, *pbUser.UserResponse) error {
	return nil
}
func (u *User) UserLogin(context.Context, *pbUser.UserRequest, *pbUser.UserResponse) error {
	return nil
}
