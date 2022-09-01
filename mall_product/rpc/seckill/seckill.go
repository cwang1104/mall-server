package rpcSeckill

import (
	"context"
	"errors"
	"fmt"
	"mall_product/common/models"
	"mall_product/common/utils"
	pbSeckill "mall_product/proto/seckill"
	"time"
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

func (s *Seckill) GetProducts(_ context.Context, in *pbSeckill.ProductRequest, out *pbSeckill.ProductResponse) error {

	products, err := models.GetProducts()

	products_rep := []*pbSeckill.Product{}

	if err != nil {
		out.Code = 500
		out.Msg = "没有查询到商品信息"
		return errors.New("没有查询到商品信息")
	}

	for _, product := range *products {
		product_rep := pbSeckill.Product{}
		product_rep.Id = int32(product.Id)
		product_rep.GoodsName = product.Name
		products_rep = append(products_rep, &product_rep)
	}

	out.Code = 200
	out.Msg = "查询成功"
	out.Products = products_rep
	return nil
}

func (s *Seckill) SecKillAdd(_ context.Context, in *pbSeckill.SecKill, out *pbSeckill.SecKillResponse) error {
	name := in.Name
	price := in.Price
	num := in.Num
	goods_id := in.GoodsID
	start_time := in.StartTime
	end_time := in.EndTime

	time_start_time, _ := time.Parse("2006-01-02 15:04:05", start_time)
	time_end_time, _ := time.Parse("2006-01-02 15:04:05", end_time)
	seckill := models.Seckills{
		Name:        name,
		Price:       price,
		Num:         int(num),
		GoodsId:     int(goods_id),
		StartTime:   time_start_time,
		EndTime:     time_end_time,
		Status:      0,
		CreatedTime: time.Now(),
	}

	err := models.AddSeckill(&seckill)

	if err != nil {
		out.Code = 500
		out.Msg = "添加失败"
	}

	out.Code = 200
	out.Msg = "添加成功"
	return nil
}

func (s *Seckill) SecKillDel(_ context.Context, in *pbSeckill.SecKillDelRequest, out *pbSeckill.SecKillResponse) error {
	id := in.Id
	// 删除数据库数据操作
	seckill := models.Seckills{
		Id: int(id),
	}
	err := models.DelSeckill(&seckill)

	if err != nil {
		out.Code = 500
		out.Msg = "删除失败"
		return nil
	}

	out.Code = 200
	out.Msg = "删除成功"
	return nil
}

func (s *Seckill) SecKillToEdit(_ context.Context, in *pbSeckill.SecKillDelRequest, out *pbSeckill.SecKillToEditResponse) error {
	id := in.Id

	seckill := models.Seckills{
		Id: int(id),
	}

	err := models.GetSeckillInfo(&seckill)
	if err != nil {
		out.Code = 500
		out.Msg = "没有查询到数据"
		return errors.New("没有查询到数据")
	}

	seckill_rep := &pbSeckill.SecKill{}
	seckill_rep.Id = int32(seckill.Id)
	seckill_rep.Name = seckill.Name
	seckill_rep.Price = seckill.Price
	seckill_rep.Num = int32(seckill.Num)
	seckill_rep.Id = int32(seckill.GoodsId)

	product, _ := models.GetProductById(seckill.GoodsId)
	seckill_rep.Name = product.Name
	seckill_rep.StartTime = seckill.StartTime.Format("2006-01-02 15:04:05")
	seckill_rep.EndTime = seckill.EndTime.Format("2006-01-02 15:04:05")

	products_no_rep := []*pbSeckill.Product{}

	products_no, _ := models.GetPordictsById(seckill.GoodsId)

	for _, product_no := range *products_no {
		product_no_rep := pbSeckill.Product{}
		product_no_rep.Id = int32(product_no.Id)
		product_no_rep.GoodsName = product_no.Name
		products_no_rep = append(products_no_rep, &product_no_rep)
	}

	out.ProductsNo = products_no_rep
	out.Code = 200
	out.Msg = "成功"
	out.Seckill = seckill_rep
	return nil
}

func (s *Seckill) SecKillDoEdit(_ context.Context, in *pbSeckill.SecKill, out *pbSeckill.SecKillResponse) error {
	id := in.Id
	name := in.Name
	price := in.Price
	num := in.Num
	Pid := in.GoodsID
	start_time := in.StartTime
	end_time := in.EndTime

	time_start_time, _ := time.Parse("2006-01-02 15:04:05", start_time)
	time_end_time, _ := time.Parse("2006-01-02 15:04:05", end_time)
	seckill := models.Seckills{
		Name:      name,
		Price:     price,
		Num:       int(num),
		GoodsId:   int(Pid),
		StartTime: time_start_time,
		EndTime:   time_end_time,
	}
	err := models.UpdateSeckill(int(id), &seckill)
	if err != nil {
		out.Code = 500
		out.Msg = "更新活动失败"
		return nil
	}
	out.Code = 200
	out.Msg = "更新活动成功"
	return nil
}

func (s *Seckill) FrontSecKillList(_ context.Context, in *pbSeckill.FrontSecKillRequest, out *pbSeckill.FrontSecKillResponse) error {
	/*
		活动显示过滤：
			1.只显示未来一天要做的活动  当前时间+1天  >=  开始时间
			2.一页6条
	*/

	tomorrow_time := utils.AddHour(24)
	fmt.Println("==============")
	fmt.Println(tomorrow_time)

	currentPage := in.CurrentPage
	pagesize := in.PageSize

	offsetNum := pagesize * (currentPage - 1)
	seckills, err := models.GetSeckillByTime(tomorrow_time, pagesize, offsetNum)
	if err != nil {
		out.Code = 500
		out.Msg = "查询不到数据"
		return errors.New("查询不到数据")
	}

	seckills_rep := []*pbSeckill.SecKill{}

	for _, seckill := range *seckills {
		seckill_rep := pbSeckill.SecKill{}
		seckill_rep.Id = int32(seckill.Id)
		seckill_rep.Name = seckill.Name
		seckill_rep.Price = seckill.Price
		seckill_rep.Num = int32(seckill.Num)
		seckill_rep.GoodsID = int32(seckill.GoodsId)

		product, _ := models.GetProductById(seckill.GoodsId)

		seckill_rep.GoodsName = product.Name
		seckill_rep.Picture = product.Picture
		seckill_rep.GoodsPrice = product.Price
		seckill_rep.GoodsDesc = product.Desc
		seckill_rep.StartTime = seckill.StartTime.Format("2006-01-02 15:04:05")
		seckill_rep.EndTime = seckill.EndTime.Format("2006-01-02 15:04:05")
		seckill_rep.CreatedTime = seckill.CreatedTime.Format("2006-01-02 15:04:05")

		seckills_rep = append(seckills_rep, &seckill_rep)
	}

	count := models.GetSecKillCountByTime(tomorrow_time)

	out.Code = 200
	out.Msg = "成功"
	out.Current = currentPage
	out.PageSize = pagesize
	out.TotalPage = (count + pagesize - 1) / pagesize
	out.SeckillList = seckills_rep

	return nil
}

func (s *Seckill) FrontSecKillDetail(_ context.Context, in *pbSeckill.SecKillDelRequest, out *pbSeckill.FrongSecKillDetailResponse) error {
	id := int(in.Id)
	fmt.Println(id)

	seckill, err := models.GetSecKillById(id)

	if err != nil {
		return errors.New("没有查询到数据")
	}
	product, _ := models.GetProductById(seckill.GoodsId)
	seckill_rep := &pbSeckill.SecKill{
		Id:          int32(seckill.Id),
		Name:        seckill.Name,
		Num:         int32(seckill.Num),
		Price:       seckill.Price,
		GoodsID:     int32(seckill.GoodsId),
		GoodsName:   product.Name,
		Picture:     product.Picture,
		GoodsPrice:  product.Price,
		GoodsDesc:   product.Desc,
		Unit:        product.Unit,
		StartTime:   seckill.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:     seckill.EndTime.Format("2006-01-02 15:04:05"),
		CreatedTime: seckill.CreatedTime.Format("2006-01-02 15:04:05"),
	}
	out.Code = 200
	out.Msg = "查询成功"
	out.Seckill = seckill_rep
	return nil
}
