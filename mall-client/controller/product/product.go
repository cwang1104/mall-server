package product

import (
	"context"
	"fmt"
	grpcc "github.com/asim/go-micro/plugins/client/grpc/v4"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"mall-client/common/utils"
	pbProduct "mall-client/proto/mall_product/product"
	"net/http"
	"strconv"
	"time"
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

func ProductAdd(c *gin.Context) {
	name := c.PostForm("name")
	price := c.PostForm("price")
	num := c.PostForm("num")
	unit := c.PostForm("unit")
	desc := c.PostForm("desc")

	pbReqParams := pbProduct.ProductAddRequest{
		Name:  name,
		Price: utils.StrToFloat32(price),
		Num:   utils.StrToInt32(num),
		Unit:  unit,
		Desc:  desc,
	}

	file, err := c.FormFile("pic")
	if err != nil {
		fmt.Println(err)
	} else {
		unix_int64 := time.Now().Unix()
		unix_str := strconv.FormatInt(unix_int64, 10)
		file_path := "upload/" + unix_str + file.Filename
		c.SaveUploadedFile(file, file_path)
		pbReqParams.Picture = file_path
	}

	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	grpcserver := pbProduct.NewProductsService("mall_product", service.Client())
	resp, err := grpcserver.ProductAdd(context.Background(), &pbReqParams)

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

func ProductDel(c *gin.Context) {
	id := c.PostForm("id")
	pbReqParams := pbProduct.ProductDelRequest{
		Id: utils.StrToInt32(id),
	}

	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	grpcserver := pbProduct.NewProductsService("mall_product", service.Client())
	resp, err := grpcserver.ProductDel(context.Background(), &pbReqParams)
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
func ProductDoEdit(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	price := c.PostForm("price")
	num := c.PostForm("num")
	unit := c.PostForm("unit")
	desc := c.PostForm("desc")

	pbReqParams := pbProduct.ProductEditRequest{
		Id:    utils.StrToInt32(id),
		Name:  name,
		Price: utils.StrToFloat32(price),
		Num:   utils.StrToInt32(num),
		Unit:  unit,
		Desc:  desc,
	}
	file, err := c.FormFile("pic")

	if err != nil {
		fmt.Println(err)
	} else {
		unix_int64 := time.Now().Unix()
		unix_str := strconv.FormatInt(unix_int64, 10)
		file_path := "upload/" + unix_str + file.Filename
		c.SaveUploadedFile(file, file_path)
		pbReqParams.Picture = file_path
	}
	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	grpcserver := pbProduct.NewProductsService("mall_product", service.Client())
	resp, err := grpcserver.ProductDoEdit(context.Background(), &pbReqParams)

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

func ProductToEdit(c *gin.Context) {
	id := c.Query("id")
	pbReqParams := pbProduct.ProductToEditRequest{
		Id: utils.StrToInt32(id),
	}

	consulReq := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(consulReq),
		micro.Client(grpcc.NewClient()),
	)

	grpcserver := pbProduct.NewProductsService("mall_product", service.Client())
	resp, err := grpcserver.ProductToEdit(context.Background(), &pbReqParams)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "grpc调用错误" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":       resp.Code,
		"msg":        resp.Msg,
		"product":    resp.Product,
		"img_base64": utils.Img2Base64(resp.Product.Picture),
	})
}


