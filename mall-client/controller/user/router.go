package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "xiaoming",
		})
	})
}
