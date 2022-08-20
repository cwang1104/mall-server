package userApi

import (
	"github.com/gin-gonic/gin"
	"mall-client/controller/middleWare"
)

func Router(router *gin.RouterGroup) {
	router.POST("/send_email", SendEmail)
	router.POST("/user_register", UserRegister)
	router.POST("/front_user_login", UserLogin)
	router.POST("/testToken", testToken)

	router.POST("/admin_login", AdminLogin)
	router.GET("/get_front_users", middleWare.ValidAdminToken, UserList)
}
