package seckill

import (
	"context"
	grpcc "github.com/asim/go-micro/plugins/client/grpc/v4"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"mall-client/common/utils"
	pbMiaoSha "mall-client/proto/mall_seckill/miaosha"
	"net/http"
)

func SeckillM(c *gin.Context) {

	id := c.PostForm("id")

	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	user_id := c.MustGet("user_id").(int)

	pbReqParams := pbMiaoSha.MiaoshaRequest{
		Id:     utils.StrToInt32(id),
		UserID: int32(user_id),
	}

	grpcserver := pbMiaoSha.NewMiaoShaService("mall_seckill", service.Client())
	resp, err := grpcserver.FrontMiaoSha(context.Background(), &pbReqParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "grpc调用错误" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})

}
