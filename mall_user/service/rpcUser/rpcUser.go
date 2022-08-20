package rpcUser

import (
	"context"
	"fmt"
	cache "github.com/patrickmn/go-cache"
	"mall_user/models"
	pbUser "mall_user/proto/user"
	"mall_user/utils"
	"time"
)

var c = cache.New(60*time.Second, 20*time.Second)

type User struct {
	pbUser.UserHandler
}

//用户注册
func (u *User) UserRegister(_ context.Context, in *pbUser.UserRequest, out *pbUser.UserResponse) error {
	//校验发送的验证码是否一致
	code := in.Code
	email := in.Email
	catch_code, is_ok := c.Get(email)
	if is_ok {
		if catch_code.(string) == code {
			//if catch_code != code {
			out.Code = 500
			out.Msg = "邮箱验证码不正确"
		} else {
			//验证邮箱是否存在，不存在则注册
			if !models.EmaiIsExist(email) {

				//md5密码加密
				md5_Password := utils.GetMd5Password(in.Password)

				user := models.User{
					Email:       email,
					Password:    md5_Password,
					Status:      1,
					Desc:        "测试用户",
					CreatedTime: time.Now(),
				}
				//注册
				err := models.RegisterUser(&user)
				if err != nil {
					out.Code = 500
					out.Msg = "注册失败"
				}
				out.Code = 200
				out.Msg = "注册成功"
			} else {
				out.Code = 500
				out.Msg = "邮箱已存在"
			}
		}
	} else {
		out.Code = 500
		out.Msg = "注册失败，请重新尝试"
	}

	return nil
}

//发送邮件
func (u *User) UserSendEmail(_ context.Context, in *pbUser.UserMailRequest, out *pbUser.UserResponse) error {
	to_email := in.Email
	randNum := utils.GetRandNum(6)
	err := utils.SendEmail(to_email, randNum)
	if err != nil {
		out.Msg = "发送邮件错误" + err.Error()
		out.Code = 500
		return nil
	}

	//缓存邮箱号及对应的随机数,便于后续登录校验
	c.Set(to_email, randNum, cache.DefaultExpiration)
	fmt.Println(c.Get(to_email))
	out.Msg = "发送邮件成功"
	out.Code = 200
	return nil
}

func (u *User) UserLogin(_ context.Context, in *pbUser.UserRequest, out *pbUser.UserResponse) error {
	email := in.Email
	password := in.Password
	if !models.EmaiIsExist(email) {
		out.Code = 500
		out.Msg = "邮箱不存在，请先注册，再登录"
		return nil
	}
	//查询数据库内容 得到user信息
	user := models.FindUserByEmail(email)
	//验证密码是否正确
	if !utils.CheckPassword(password, user.Password) { //密码错误
		out.Code = 500
		out.Msg = "密码错误"
		return nil
	} else { //密码正确
		out.Code = 200
		out.Msg = "登录成功"
		out.Email = user.Email
	}
	return nil
}
