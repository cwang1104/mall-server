package userApi

import (
	"context"
	"fmt"
	grpcc "github.com/asim/go-micro/plugins/client/grpc/v4"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"mall-client/common/utils"
	pbUser "mall-client/proto/mall_user/user"
	"net/http"
)

func SendEmail(c *gin.Context) {
	email := c.PostForm("email")

	//验证邮箱格式，不是则返回
	if utils.VerifyEmail(email) == false {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "邮箱格式不正确~",
		})
		return
	}
	//邮箱格式正确,发送调用user服务发送邮件
	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	pbReqParams := pbUser.UserMailRequest{
		Email: email,
	}

	grpcserver := pbUser.NewUserService("mall_user", service.Client())
	resp, err := grpcserver.UserSendEmail(context.TODO(), &pbReqParams)
	if err != nil {
		fmt.Println("get grpc userSendEmail err", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "调用grpc错误~",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}
