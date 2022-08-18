package userApi

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.POST("/send_email", SendEmail)
}
