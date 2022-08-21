package product

import (
	"context"
	grpcc "github.com/asim/go-micro/plugins/client/grpc/v4"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"mall-client/common/utils"
	pbProduct "mall-client/proto/mall_product/product"
	"net/http"
)

//type GetPorductListReq struct {
//	CurrentPage int32 `uri:"currentPage"`
//	PageSize    int32 `uri:"pageSize"`
//}

func GetProductList(c *gin.Context) {
	//var req GetPorductListReq
	//if err := c.ShouldBindUri(&req); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"code": 500,
	//		"msg":  "参数获取错误",
	//	})
	//	return
	//}
	currentPage := c.DefaultQuery("currentPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	// grpc调用
	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	pbReqParams := pbProduct.ProductsRequest{
		PageSize:    utils.StrToInt32(pageSize),
		CurrentPage: utils.StrToInt32(currentPage),
	}

	grpcserver := pbProduct.NewProductsService("mall_product", service.Client())
	resp, err := grpcserver.ProductList(context.Background(), &pbReqParams)
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
			"products":     resp.Products,
			"total":        resp.Total,
			"current_page": resp.Current,
			"page_size":    resp.PageSize,
		})
	}
}

func ProductDoEdit(c *gin.Context) {

}

func ProductToEdit(c *gin.Context) {

}

func ProductDel(c *gin.Context) {

}

func ProductAdd(c *gin.Context) {

}
