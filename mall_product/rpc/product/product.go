package rpcProduct

import (
	"context"
	"log"
	"mall_product/common/models"
	pbProduct "mall_product/proto/product"
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
