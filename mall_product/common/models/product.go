package models

import "fmt"

func GetProductLists(pageSize, CurrentPage int32) (*[]Product, error) {
	offsetNum := pageSize * (CurrentPage - 1)
	var product []Product
	result := db.Limit(pageSize).Offset(offsetNum).Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(product)
	return &product, nil
}

func GetProductCount() (*int, error) {
	var count int
	result := db.Find(&[]Product{}).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	return &count, nil
}

func AddProduct(product *Product) error {
	result := db.Create(product)
	return result.Error
}

func DelProduct(product *Product) error {
	result := db.Delete(product)
	return result.Error
}

func GetProductById(id int) (*Product, error) {
	var product Product
	result := db.Where("id = ?", id).First(&product)
	return &product, result.Error
}

func GetProduct(product *Product) {
	db.First(product)
}

func UpdateProduct(product *Product, id int32) error {
	return db.Where("id = ?", int(id)).Find(&Product{}).Update(product).Error
}

func GetProducts() (*[]Product, error) {
	var products []Product
	result := db.Find(&products)
	return &products, result.Error
}

func GetPordictsById(id int) (*[]Product, error) {
	products_no := []Product{}
	result := db.Where("id != ?", id).Find(&products_no)
	return &products_no, result.Error
}
