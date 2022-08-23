package rpcSeckill

import (
	"context"
	"mall_product/common/models"
	pbSeckill "mall_product/proto/seckill"
)

type Seckill struct {
	pbSeckill.SecKillsHandler
}

func (s *Seckill) SecKillList(_ context.Context, in *pbSeckill.SecKillsRequest, out *pbSeckill.SecKillsResponse) error {
	currentPage := in.CurrentPage
	pagesize := in.PageSize

	offsetNum := pagesize * (currentPage - 1)

	seckills, err := models.GetSecKillList(int(pagesize), int(offsetNum))

	if err != nil {
		out.Code = 500
		out.Msg = "没有查询到数据 " + err.Error()
		return nil
	}

	count, err := models.GetSecKillCount()
	if err != nil {
		out.Code = 500
		out.Msg = "没有查询到数据 " + err.Error()
		return nil
	}

	seckills_resp := []*pbSeckill.SecKill{}
	for _, seckill := range *seckills {
		seckill_resp := pbSeckill.SecKill{}
		seckill_resp.Id = int32(seckill.Id)
		seckill_resp.Name = seckill.Name
		seckill_resp.Price = seckill.Price
		seckill_resp.Num = int32(seckill.Num)
		product := models.Product{
			Id: seckill.GoodsId,
		}
		models.GetProduct(&product)
		seckill_resp.GoodsID = int32(seckill.GoodsId)
		seckill_resp.GoodsName = product.Name
		seckill_resp.StartTime = seckill.StartTime.Format("2006-01-02 15:04:05")
		seckill_resp.EndTime = seckill.EndTime.Format("2006-01-02 15:04:05")
		seckill_resp.CreatedTime = seckill.CreatedTime.Format("2006-01-02 15:04:05")

		seckills_resp = append(seckills_resp, &seckill_resp)
	}
	out.Code = 200
	out.Msg = "查询成功"
	out.Current = currentPage
	out.PageSize = pagesize
	out.Seckills = seckills_resp
	out.Total = int32(count)
	return nil
}
