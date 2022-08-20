package rpcAdmin

import (
	"context"
	"mall_user/models"
	pbAdmin "mall_user/proto/admin"
	"mall_user/utils"
	"strconv"
)

type Admin struct {
	pbAdmin.AdminHandler
}

func (a *Admin) AdminLogin(_ context.Context, in *pbAdmin.AdminRequest, out *pbAdmin.AdminResponse) error {

	user_name := in.UserName
	password := in.Password
	if !models.IsAdminExist(user_name) {
		out.Code = 500
		out.Msg = "用户名"
		return nil
	}
	//查询数据库内容 得到user信息
	admin := models.FindAdminByUserName(user_name)
	//验证密码是否正确
	if !utils.CheckPassword(password, admin.Password) { //密码错误
		out.Code = 500
		out.Msg = "密码错误"
		return nil
	} else { //密码正确
		out.Code = 200
		out.Msg = "登录成功"
		out.UserName = admin.UserName
	}
	return nil
}

func (a *Admin) GetUserList(_ context.Context, in *pbAdmin.UserRequest, out *pbAdmin.UserResponse) error {
	currentPage := int(in.CurrentPage)
	pageSize := int(in.PageSize)

	/*
		current offset limit
		1       0        2       2 * (1 - 1)
		2       2        2		 2 * (2 - 1)
		3       4         2		2 * (3 -1 )

		offset = limit * (current - 1)
	*/

	offsetNum := pageSize * (currentPage - 1)

	userLists, err := models.GetUserLists(pageSize, offsetNum)
	if err != nil {
		out.Code = 500
		out.Msg = "没有查询到数据"
	}

	UsersCount := models.GetUserCount()

	//将数据库格式信息转换为grpc定义的数据类型并返回
	pbUsers := []*pbAdmin.User{}
	pb_user := pbAdmin.User{}
	for _, user := range *userLists {
		pb_user.Email = user.Email
		pb_user.Status = strconv.Itoa(user.Status)
		pb_user.Desc = user.Desc
		pb_user.CreatedTime = user.CreatedTime.String()
		pbUsers = append(pbUsers, &pb_user)
	}
	out.Code = 200
	out.Msg = "成功"
	out.Users = pbUsers
	out.Total = int32(*UsersCount)
	out.Current = int32(currentPage)
	out.PageSize = int32(pageSize)
	return nil

}
