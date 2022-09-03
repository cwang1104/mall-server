package middleWare

import (
	"github.com/gin-gonic/gin"
	"mall-client/common/utils"
	"net/http"
	"strings"
)

func ValidAdminToken(c *gin.Context) {

	au_header := c.Request.Header.Get("Authorization")
	if len(au_header) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请携带token",
		})
		c.Abort()
		return
	}

	//将Header进行解析为授权头和token
	tokenFields := strings.Fields(au_header)

	tokenType := strings.ToLower(tokenFields[0]) //转为小写方便比较
	if tokenType != "bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": "500",
			"msg":  "token type is not bearer",
		})
		c.Abort()
		return
	}

	accessToken := tokenFields[1]

	claims, err := utils.AuthToken(accessToken, utils.AdminSecretKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": "500",
			"msg":  "auth token err" + err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("username", claims.UserName)
	c.Next()
}

func ValidUserToken(c *gin.Context) {
	au_header := c.Request.Header.Get("Authorization")
	if len(au_header) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请携带token",
		})
		c.Abort()
		return
	}

	//将Header进行解析为授权头和token
	tokenFields := strings.Fields(au_header)

	tokenType := strings.ToLower(tokenFields[0]) //转为小写方便比较
	if tokenType != "bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": "500",
			"msg":  "token type is not bearer",
		})
		c.Abort()
		return
	}

	accessToken := tokenFields[1]

	claims, err := utils.AuthToken(accessToken, utils.UserSecretKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": "500",
			"msg":  "auth token err" + err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("user_id", claims.UserId)
	c.Next()
}
