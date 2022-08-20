package userApi

import (
	"context"
	grpcc "github.com/asim/go-micro/plugins/client/grpc/v4"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"mall-client/common/utils"
	pbAdmin "mall-client/proto/mall_user/admin"
	"net/http"
)

type AdminReqParams struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

func AdminLogin(c *gin.Context) {
	var req AdminReqParams
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "获取参数错误" + err.Error(),
		})
		return
	}

	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	pbReqParams := pbAdmin.AdminRequest{
		UserName: req.UserName,
		Password: req.Password,
	}

	grpcserver := pbAdmin.NewAdminService("mall_user", service.Client())
	resp, err := grpcserver.AdminLogin(context.Background(), &pbReqParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "grpc调用失败" + err.Error(),
		})
		return
	}
	if resp.Code != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  resp.Msg,
		})
		return
	}

	admin_token, err1 := utils.GenToken(resp.UserName, utils.AdminExpireDuration, utils.AdminSecretKey)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "token错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":        resp.Code,
		"msg":         resp.Msg,
		"admin_token": admin_token,
		"user_name":   resp.UserName,
	})

}

func UserList(c *gin.Context) {
	currentPage := c.DefaultQuery("currentPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	pbReqParams := pbAdmin.UserRequest{
		CurrentPage: utils.StrToInt32(currentPage),
		PageSize:    utils.StrToInt32(pageSize),
	}

	grpcserver := pbAdmin.NewAdminService("mall_user", service.Client())
	rep, err := grpcserver.GetUserList(context.Background(), &pbReqParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "grpc调用失败" + err.Error(),
		})
		return
	}
	if rep.Code != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":         200,
		"msg":          "成功",
		"front_users":  rep.Users,
		"total":        rep.Total,
		"current_page": rep.Current,
		"page_size":    rep.PageSize,
	})

}
