package models

import "github.com/jinzhu/gorm"

func AddOrderInfo(order *Order) error {
	return db.Table("order").Create(order).Error
}

func OrderExist(user_id, ac_id int) error {
	return db.Where("user_id = ?", user_id).Where("activity_id = ?", ac_id).Find(&Order{}).Error
}

func OrderResult(user_id, ac_id int) *gorm.DB {
	return db.Where("user_id = ?", user_id).Where("activity_id = ?", ac_id).Find(&Order{})
}
