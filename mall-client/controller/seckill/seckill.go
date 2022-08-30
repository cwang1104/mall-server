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

func GetProducts(c *gin.Context) {
	// 返回商品name 和 id

	// 和srv通信获取products数据
	// grpc调用
	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	pbReqParams := pbSeckill.ProductRequest{}

	grpcserver := pbSeckill.NewSecKillsService("mall_product", service.Client())
	resp, err := grpcserver.GetProducts(context.Background(), &pbReqParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "grpc调用错误" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"msg":      resp.Msg,
		"products": resp.Products,
	})
}

func SecKillAdd(c *gin.Context) {
	name := c.PostForm("name")
	price := c.PostForm("price")
	num := c.PostForm("num")
	pid := c.PostForm("pid")
	start_time := c.PostForm("start_time")
	end_time := c.PostForm("end_time")

	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	pbReqParams := pbSeckill.SecKill{
		Name:      name,
		Price:     utils.StrToFloat32(price),
		Num:       utils.StrToInt32(num),
		GoodsID:   utils.StrToInt32(pid),
		StartTime: start_time,
		EndTime:   end_time,
	}

	grpcserver := pbSeckill.NewSecKillsService("mall_product", service.Client())
	resp, err := grpcserver.SecKillAdd(context.Background(), &pbReqParams)
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

func SecKillDel(c *gin.Context) {
	id := c.PostForm("id")
	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)
	pbReqParams := pbSeckill.SecKillDelRequest{
		Id: utils.StrToInt32(id),
	}

	grpcserver := pbSeckill.NewSecKillsService("mall_product", service.Client())
	resp, err := grpcserver.SecKillDel(context.Background(), &pbReqParams)
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

func SeckillToEdit(c *gin.Context) {
	id := c.Query("id")

	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)
	pbReqParams := pbSeckill.SecKillDelRequest{
		Id: utils.StrToInt32(id),
	}
	grpcserver := pbSeckill.NewSecKillsService("mall_product", service.Client())
	resp, err := grpcserver.SecKillToEdit(context.Background(), &pbReqParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "grpc调用错误" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":        resp.Code,
		"msg":         resp.Msg,
		"seckill":     resp.Seckill,
		"products_no": resp.ProductsNo,
	})
}

func ProductDoEdit(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	price := c.PostForm("price")
	num := c.PostForm("num")
	pid := c.PostForm("pid")
	start_time := c.PostForm("start_time")
	end_time := c.PostForm("end_time")

	// 和srv通信获取front_users数据
	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)
	pbReqParams := pbSeckill.SecKill{
		Id:        utils.StrToInt32(id),
		Name:      name,
		Num:       utils.StrToInt32(num),
		Price:     utils.StrToFloat32(price),
		StartTime: start_time,
		EndTime:   end_time,
		GoodsID:   utils.StrToInt32(pid),
	}
	grpcserver := pbSeckill.NewSecKillsService("mall_product", service.Client())
	resp, err := grpcserver.SecKillDoEdit(context.Background(), &pbReqParams)
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

func GetFrontSeckillList(c *gin.Context) {

	currentPage := c.DefaultQuery("currentPage", "1")
	pageSize := c.DefaultQuery("pageSize", "8")

	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)
	pbReqParams := pbSeckill.FrontSecKillRequest{
		CurrentPage: utils.StrToInt32(currentPage),
		PageSize:    utils.StrToInt32(pageSize),
	}
	grpcserver := pbSeckill.NewSecKillsService("mall_product", service.Client())
	resp, err := grpcserver.FrontSecKillList(context.Background(), &pbReqParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "grpc调用错误" + err.Error(),
		})
		return
	}

	for _, seckill := range resp.SeckillList {
		seckill.Picture = utils.Img2Base64(seckill.Picture)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":         resp.Code,
		"msg":          resp.Msg,
		"current":      resp.Current,
		"page_size":    resp.PageSize,
		"total_page":   resp.TotalPage,
		"seckill_list": resp.SeckillList,
	})

}

func SecKillDetail(c *gin.Context) {
	id := c.Query("id")

	// grpc 通信
	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)
	pbReqParams := pbSeckill.SecKillDelRequest{
		Id: utils.StrToInt32(id),
	}
	grpcserver := pbSeckill.NewSecKillsService("mall_product", service.Client())
	resp, err := grpcserver.FrontSecKillDetail(context.Background(), &pbReqParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "grpc调用错误" + err.Error(),
		})
		return
	}

	resp.Seckill.Picture = utils.Img2Base64(resp.Seckill.Picture)
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"msg":     resp.Msg,
		"seckill": resp.Seckill,
	})
}
