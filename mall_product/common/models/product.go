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
