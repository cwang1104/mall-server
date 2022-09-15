package seckill

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"mall-client/common/utils"
	"mall-client/rabbitmq"
	"net/http"
	"strconv"
)

func SeckillM(c *gin.Context) {

	id := c.PostForm("id")
	user_id := c.MustGet("user_id").(int)

	orderMap := map[string]interface{}{
		"user_id": strconv.Itoa(user_id),
		"id":      id,
	}

	str, err := json.Marshal(orderMap)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "系统错误",
		})
		return
	}

	qe := rabbitmq.QueueAndExchange{
		QueueName:    "mall.web.Queue",
		ExchangeName: "mall.web.exchange",
		ExchangeType: "direct",
		RoutingKey:   "mall.web.routingKey",
	}

	mq := rabbitmq.NewRabbitMq(qe)
	mq.ConnMq()
	mq.OpenChan()

	defer func() {
		mq.CloseMq()
	}()

	defer func() {
		mq.CloseChan()
	}()

	mq.PublishMsg(string(str))

	c.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  "抢购进行中",
	})

}

func GetSeckillResult(c *gin.Context) {
	id, exist := c.Get("user_id")

	fmt.Println("==============")
	fmt.Println(exist)
	if exist {
		conn, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			fmt.Println("连接出错")
		}

		ret, err_r := redis.String(conn.Do("get", id))
		fmt.Println(ret)
		fmt.Println(err_r)
		if err_r == nil {
			ret_map := utils.StrToMap(ret)
			fmt.Println(ret_map)
			c.JSON(http.StatusOK, gin.H{ // 说明从redis里面获取到了数据，
				"code": 200,
				"msg":  ret_map["msg"],
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 500,
		})
		return
	}

}
