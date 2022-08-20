package userApi

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.POST("/send_email", SendEmail)
	router.POST("/user_register", UserRegister)
	router.POST("/login", UserLogin)
	router.POST("/testToken", testToken)
}
