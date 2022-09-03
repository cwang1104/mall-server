package rpcMiaoSha

import (
	"context"
	"mall_seckill/common/models"
	pbMiaoSha "mall_seckill/proto/miaosha"
	"time"
)

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
