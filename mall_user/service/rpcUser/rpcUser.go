package rpcUser

import (
	"context"
	"fmt"
	cache "github.com/patrickmn/go-cache"
	pbUser "mall_user/proto/user"
	"mall_user/utils"
	"time"
)

type User struct {
	pbUser.UserHandler
}

//用户注册
func (u *User) UserRegister(ctx context.Context, in *pbUser.UserRequest, out *pbUser.UserResponse) error {

	return nil
}

//发送邮件
func (u *User) UserSendEmail(ctx context.Context, in *pbUser.UserMailRequest, out *pbUser.UserResponse) error {
	to_email := in.Email
	randNum := utils.GetRandNum(6)
	err := utils.SendEmail(to_email, randNum)
	if err != nil {
		out.Msg = "发送邮件错误"
		out.Code = 500
		return err
	}

	//缓存邮箱号及对应的随机数,便于后续登录校验
	c := cache.New(180*time.Second, 60*time.Second)
	c.Set(to_email, randNum, cache.DefaultExpiration)
	fmt.Println(c.Get(to_email))
	out.Msg = "发送邮件成功"
	out.Code = 200
	return nil
}
func (u *User) UserLogin(ctx context.Context, in *pbUser.UserRequest, out *pbUser.UserResponse) error {
	return nil
}
