package seckill

import (
	"github.com/gin-gonic/gin"
	"mall-client/controller/middleWare"
)

func Router(router *gin.RouterGroup) {
	// 管理端
	router.GET("/get_seckill_list", middleWare.ValidAdminToken, GetSeckillList)
}
