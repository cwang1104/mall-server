package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mall-client/controller"
)

func main() {

	router := gin.Default()

	//解决跨域请求
	router.Use(cors.Default())
	controller.InitRouter(router)
	err := router.Run("127.0.0.1:9999")
	if err != nil {
		fmt.Println("run err", err)
		return
	}
}
