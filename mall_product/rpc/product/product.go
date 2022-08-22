package rpcProduct

import (
	"context"
	"log"
	"mall_product/common/models"
	pbProduct "mall_product/proto/product"
	"time"
)

type Product struct {
	pbProduct.ProductsHandler
}

func (p *Product) ProductList(_ context.Context, in *pbProduct.ProductsRequest, out *pbProduct.ProductsResponse) error {
	currentPage := in.CurrentPage
	pagesize := in.PageSize

	products, err := models.GetProductLists(pagesize, currentPage)
	if err != nil {
		out.Code = 500
		out.Msg = "没有查询到数据" + err.Error()
		return nil
	}

	product_count, err := models.GetProductCount()
	if err != nil {
		out.Code = 500
		out.Msg = "没有查询到数据" + err.Error()
		return nil
	}
	products_resp := []*pbProduct.Product{}

	for _, product := range *products {
		product_resp := pbProduct.Product{}
		product_resp.Id = int32(product.Id)
		product_resp.Name = product.Name
		product_resp.Price = product.Price
		product_resp.Num = int32(product.Num)
		product_resp.Unit = product.Unit
		product_resp.Picture = product.Picture
		product_resp.Desc = product.Desc
		product_resp.CreateTime = product.CreatedTime.String()

		products_resp = append(products_resp, &product_resp)
	}

	log.Println(products_resp)

	out.Code = 200
	out.Msg = "查询成功"
	out.Current = currentPage
	out.PageSize = pagesize
	out.Products = products_resp
	out.Total = int32(*product_count)

	return nil
}

func (p *Product) ProductAdd(_ context.Context, in *pbProduct.ProductAddRequest, out *pbProduct.ProductAddResponse) error {
	name := in.Name
	price := in.Price
	num := in.Num
	unit := in.Unit
	pic_path := in.Picture
	desc := in.Desc

	product := models.Product{
		Name:        name,
		Price:       price,
		Num:         int(num),
		Unit:        unit,
		Desc:        desc,
		Picture:     pic_path,
		CreatedTime: time.Now(),
	}
	err := models.AddProduct(&product)

	if err != nil {
		out.Code = 500
		out.Msg = "添加商品失败"
		return nil
	}

	out.Code = 200
	out.Msg = "添加商品成功"
	return nil
}

func (p *Product) ProductDel(_ context.Context, in *pbProduct.ProductDelRequest, out *pbProduct.ProductAddResponse) error {
	id := in.Id
	product := models.Product{
		Id: int(id),
	}
	err := models.DelProduct(&product)
	if err != nil {
		out.Code = 500
		out.Msg = "删除失败"
		return nil
	}
	out.Code = 200
	out.Msg = "删除成功"
	return nil
}

func (p *Product) ProductToEdit(_ context.Context, in *pbProduct.ProductToEditRequest, out *pbProduct.ProductToEditResponse) error {
	id := in.Id
	product, err := models.GetProductById(int(id))
	if err != nil {
		out.Code = 500
		out.Msg = "数据库查询出错" + err.Error()
		return nil
	}
	product_rep := &pbProduct.Product{}
	product_rep.Id = int32(product.Id)
	product_rep.Name = product.Name
	product_rep.Price = product.Price
	product_rep.Num = int32(product.Num)
	product_rep.Unit = product.Unit
	product_rep.Picture = product.Picture
	product_rep.Desc = product.Desc

	out.Code = 200
	out.Msg = "查询成功"
	out.Product = product_rep
	return nil
}
