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
	"strings"
)

//发送邮件
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

type userRegisterParams struct {
	Email      string `form:"email"`
	Catpche    string `form:"catpche"`
	Password   string `form:"password"`
	Repassword string `form:"repassword"`
}

// 注册相关
func UserRegister(c *gin.Context) {
	var registerReq userRegisterParams
	if err := c.ShouldBind(&registerReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "200",
			"msg":  "参数获取失败~",
		})
		return
	}

	//校验邮箱
	isOk := utils.VerifyEmail(registerReq.Email)
	if !isOk {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "200",
			"msg":  "邮箱格式不正确~",
		})
		return
	}

	//校验两种密码是否一致
	if registerReq.Password != registerReq.Repassword {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "200",
			"msg":  "两次密码不匹配~",
		})
		return
	}

	//grpc调用完成注册
	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	pbReqParams := pbUser.UserRequest{
		Email:     registerReq.Email,
		Code:      registerReq.Catpche,
		Password:  registerReq.Password,
		Reassword: registerReq.Repassword,
	}

	grpcserver := pbUser.NewUserService("mall_user", service.Client())

	resp, err := grpcserver.UserRegister(context.TODO(), &pbReqParams)
	if err != nil {
		fmt.Println("get grpc userSendEmail err", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "get grpc userSendEmail err",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})

}

type userLoginParams struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

//用户登录
func UserLogin(c *gin.Context) {
	var req userLoginParams
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "500",
			"msg":  "参数获取失败~",
		})
		return
	}

	//todo 访问user服务 查询数据库进行信息校验
	//grpc调用完成注册
	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	pbReqParams := pbUser.UserRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	grpcserver := pbUser.NewUserService("mall_user", service.Client())
	resp, err := grpcserver.UserLogin(context.Background(), &pbReqParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "grpc server error",
		})
		return
	}
	if resp.Code != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
		return
	}

	//用户服务端验证通过，生成token返回前端
	token, err := utils.GenToken(req.Email, utils.UserExpireDuration, utils.UserSecretKey)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "500",
			"msg":  "token获取失败~",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"msg":      resp.Msg,
		"token":    token,
		"username": resp.Email,
	})
	return
}

func testToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  "500",
			"msg":   "token if not fond",
			"token": token,
		})
		return
	}

	//将Header进行解析为授权头和token
	tokenFields := strings.Fields(token)

	tokenType := strings.ToLower(tokenFields[0]) //转为小写方便比较
	if tokenType != "bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": "500",
			"msg":  "token type is not bearer",
		})
		return
	}

	accessToken := tokenFields[1]

	claims, err := utils.AuthToken(accessToken, utils.UserSecretKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  "500",
			"msg":   "auth token err",
			"token": token,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "ok",
		"data":  claims,
		"token": token,
	})
}
