package seckill

import (
	"context"
	grpcc "github.com/asim/go-micro/plugins/client/grpc/v4"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"mall-client/common/utils"
	pbSeckill "mall-client/proto/mall_product/seckill"
	"net/http"
)

func GetSeckillList(c *gin.Context) {
	currentPage := c.DefaultQuery("currentPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	// grpc调用
	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	pbReqParams := pbSeckill.SecKillsRequest{
		PageSize:    utils.StrToInt32(pageSize),
		CurrentPage: utils.StrToInt32(currentPage),
	}

	grpcserver := pbSeckill.NewSecKillsService("mall_product", service.Client())
	resp, err := grpcserver.SecKillList(context.Background(), &pbReqParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "grpc调用错误" + err.Error(),
		})
		return
	}
	if resp.Code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":         resp.Code,
			"msg":          resp.Msg,
			"seckills":     resp.Seckills,
			"total":        resp.Total,
			"current_page": resp.Current,
			"page_size":    resp.PageSize,
		})
	}
}
