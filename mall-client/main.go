package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall-client/controller"
	"mall-client/controller/middleWare"
)

func main() {

	router := gin.Default()

	//解决跨域请求
	//router.Use(cors.Default())

	router.Use(middleWare.CrosMiddleWare)
	controller.InitRouter(router)
	err := router.Run("127.0.0.1:9999")

	if err != nil {
		fmt.Println("run err", err)
		return
	}
}
