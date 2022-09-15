package seckill

import (
	"github.com/gin-gonic/gin"
	"mall-client/controller/middleWare"
)

func Router(router *gin.RouterGroup) {
	// 管理端
	router.GET("/get_seckill_list", middleWare.ValidAdminToken, GetSeckillList)
	router.GET("/get_products", middleWare.ValidAdminToken, GetProducts)
	router.POST("/seckill_add", middleWare.ValidAdminToken, SecKillAdd)
	router.POST("/seckill_del", middleWare.ValidAdminToken, SecKillDel)
	router.GET("/seckill_to_edit", middleWare.ValidAdminToken, SeckillToEdit)
	router.POST("/seckill_do_edit", middleWare.ValidAdminToken, ProductDoEdit)

	// 前端列表
	router.GET("/front/get_seckill_list", GetFrontSeckillList)
	// 前端详情
	router.GET("/front/seckill_detail", middleWare.ValidUserToken, SecKillDetail)

	//秒杀接口
	router.POST("/front/seckill", middleWare.ValidUserToken, SeckillM)
	router.GET("/front/get_seckill_result", middleWare.ValidUserToken, GetSeckillResult)
}
