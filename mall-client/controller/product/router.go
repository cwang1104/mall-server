package product

import (
	"github.com/gin-gonic/gin"
	"mall-client/controller/middleWare"
)

func Router(router *gin.RouterGroup) {
	router.GET("/get_product_list", GetProductList)
	router.POST("/product_add", middleWare.ValidAdminToken, ProductAdd)
	router.POST("/product_del", middleWare.ValidAdminToken, ProductDel)
	router.GET("/to_product_edit", middleWare.ValidAdminToken, ProductToEdit)
	router.POST("/do_product_edit", middleWare.ValidAdminToken, ProductDoEdit)
}
