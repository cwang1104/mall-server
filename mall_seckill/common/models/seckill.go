package models

import "github.com/jinzhu/gorm"

func GetSeckillById(id int32) (*gorm.DB, *Seckills) {
	var seckill Seckills
	result := db.Where("id = ?", id).Find(&seckill)
	return result, &seckill
}

func UpdateSeckillNum(seckills *Seckills) error {
	return db.Find(seckills).Update("num", seckills.Num-1).Error
}
