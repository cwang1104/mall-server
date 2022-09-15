package rpcMiaoSha

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"mall_seckill/common/models"
	"mall_seckill/common/redis"
	"mall_seckill/common/utils"
	pbMiaoSha "mall_seckill/proto/miaosha"
	"strconv"
	"time"
)

const urls = "amqp://test01:123456@127.0.0.1:5672//test"

type MiaoSha struct {
	pbMiaoSha.MiaoShaHandler
}

func (s *MiaoSha) FrontMiaoSha(_ context.Context, in *pbMiaoSha.MiaoshaRequest, out *pbMiaoSha.MiaoShaResponse) error {

	/*
		秒杀逻辑：
			减库存

		限制：
			1、生成订单
			2、活动的数量为0不能给i有奶
			3、每个用户只能购买一个：第一次生成订单-查询订单是否存在
			4、开始时间和结束时间校验，未开始和已结束不能购买
			5、
	*/

	//查询活动信息
	result, seckill := models.GetSeckillById(in.Id)
	if result.Error != nil {
		out.Code = 500
		out.Msg = "下单失败" + result.Error.Error()
		return nil
	}

	//时间校验
	now_time := time.Now().Unix()
	if seckill.StartTime.Unix() >= now_time || seckill.EndTime.Unix() < now_time {
		out.Code = 500
		out.Msg = "下单失败-- time err"
		return nil
	}

	//库存大于0 才可继续抢购
	num_result := result.Where("num > 0").Find(seckill)
	if num_result.Error != nil {
		out.Code = 500
		out.Msg = "下单失败,库存不足" + num_result.Error.Error()
		return nil
	}

	//先查询是否有订单
	user_id := in.UserID
	err := models.OrderExist(int(user_id), int(in.Id))
	if err == nil { //查询有错，则证明没有数据，可以下单
		out.Code = 500
		out.Msg = "下单失败，订单已经存在"
		return nil
	}

	//扣减库存
	//err = models.UpdateSeckillNum(miaosha)
	ret := result.Update("num", seckill.Num-1)
	if ret.Error != nil {
		out.Code = 500
		out.Msg = "下单失败" + ret.Error.Error()
		return nil
	}

	//扣减库存后插入订单表
	order := models.Order{
		OrderNum:    1,
		UserId:      int(user_id),
		ActivityId:  int(in.Id),
		PayStatus:   1,
		CreatedTime: time.Now(),
	}
	err = models.AddOrderInfo(&order)
	if err != nil {
		out.Code = 500
		out.Msg = "创建订单失败" + err.Error()
		return nil
	}

	out.Code = 200
	out.Msg = "下单成功"
	return nil
}

func init() {
	conn, err := amqp.Dial(urls)
	if err != nil {
		fmt.Println("dial amqp conn", err)
	}

	ch, err_ch := conn.Channel()
	fmt.Println(err_ch)
	defer ch.Close()

	err = ch.Qos(1, 0, false)
	fmt.Println(err)

	deliveries, err := ch.Consume("mall.web.Queue", "order_consumer", false, false, false, false, nil)
	fmt.Println(err)
	for delivery := range deliveries {
		//fmt.Println(delivery.ContentType)
		//fmt.Println(string(delivery.Body))
		//delivery.Ack(true)
		fmt.Println("接收到任务")
		go Orderapply(delivery)
	}
}

func Orderapply(delivery amqp.Delivery) {
	body := delivery.Body

	var reqData map[string]interface{}
	_ = json.Unmarshal(body, &reqData)

	id, _ := strconv.Atoi(reqData["id"].(string))
	userId, _ := strconv.Atoi(reqData["user_id"].(string))

	//避免重复消费
	err := models.OrderExist(userId, id) //err== nil 存在
	if err == nil {
		delivery.Ack(true)
		return
	}

	now_time := time.Now().Unix()
	result, seckill := models.GetSeckillById(int32(id))
	if result.Error != nil {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "下单失败",
		}
		is_ok, _ := redis.RedisConn.Do("GET", userId)
		if is_ok == "ok" {
			return
		}

		redis.RedisConn.Do("SET", userId, utils.MapToStr(map_data))
		return
	}

	if seckill.StartTime.Unix() >= now_time {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "抢购还未开始",
		}
		is_ok, _ := redis.RedisConn.Do("GET", userId)
		if is_ok == "ok" {
			return
		}

		redis.RedisConn.Do("SET", userId, utils.MapToStr(map_data))
		return

	}

	if seckill.EndTime.Unix() < now_time {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "抢购已经结束",
		}
		is_ok, _ := redis.RedisConn.Do("GET", userId)
		if is_ok == "ok" {
			return
		}

		redis.RedisConn.Do("SET", userId, utils.MapToStr(map_data))
		return
	}

	num_result := result.Where("num > 0").Find(seckill)
	if num_result.Error != nil {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "已经抢完",
		}
		is_ok, _ := redis.RedisConn.Do("GET", userId)
		if is_ok == "ok" {
			return
		}

		redis.RedisConn.Do("SET", userId, utils.MapToStr(map_data))
		return
	}

	//获取用户信息

	ret := models.OrderResult(userId, id)
	if ret.Error == nil {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "只能买一个",
		}
		is_ok, _ := redis.RedisConn.Do("GET", userId)
		if is_ok == "ok" {
			return
		}

		redis.RedisConn.Do("SET", userId, utils.MapToStr(map_data))
		return
	}

	ret_up := result.Update("num", seckill.Num-1)

	//生成订单
	order := models.Order{
		OrderNum:    1,
		UserId:      userId,
		ActivityId:  id,
		PayStatus:   1,
		CreatedTime: time.Now(),
	}
	err = models.AddOrderInfo(&order)
	if err != nil {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "创建订单失败",
		}
		is_ok, _ := redis.RedisConn.Do("GET", userId)
		if is_ok == "ok" {
			return
		}

		redis.RedisConn.Do("SET", userId, utils.MapToStr(map_data))
		return
	}

	if ret_up.Error != nil {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "扣减库存失败",
		}
		is_ok, _ := redis.RedisConn.Do("GET", userId)
		if is_ok == "ok" {
			return
		}

		redis.RedisConn.Do("SET", userId, utils.MapToStr(map_data))
		return
	}

	//成功
	delivery.Ack(true)
	map_data := map[string]interface{}{
		"code": 200,
		"msg":  "下单成功",
	}
	is_ok, _ := redis.RedisConn.Do("GET", userId)
	if is_ok == "ok" {
		return
	}

	redis.RedisConn.Do("SET", userId, utils.MapToStr(map_data))
	return

}
