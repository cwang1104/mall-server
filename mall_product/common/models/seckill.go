package models

func GetSecKillList(pagesize, offsetNum int) (*[]Seckills, error) {
	var seckills []Seckills
	result := db.Limit(pagesize).Offset(offsetNum).Find(&seckills)
	return &seckills, result.Error
}

func GetSecKillCount() (int, error) {
	var count int
	result := db.Find(&[]Seckills{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func AddSeckill(seckill *Seckills) error {
	return db.Table("seckills").Create(seckill).Error
}

func DelSeckill(seckill *Seckills) error {
	return db.Table("seckills").Delete(seckill).Error

}

func GetSeckillInfo(seckill *Seckills) error {
	return db.Table("seckills").First(seckill).Error
}

func UpdateSeckill(id int, kills *Seckills) error {
	return db.Where("id = ?", id).Find(&Seckills{}).Update(kills).Error

}

func GetSeckillByTime(tomorrow_time string, page_size, offset int32) (*[]Seckills, error) {
	var seckills []Seckills
	result := db.Where("start_time <= ?", tomorrow_time).Where("status = ?", 0).Limit(page_size).Offset(offset).Find(&seckills)
	return &seckills, result.Error
}

func GetSecKillCountByTime(tomorrow_time string) int32 {
	seckills_count := []Seckills{}
	var count int32
	db.Where("start_time <= ?", tomorrow_time).Where("status = ?", 0).Find(&seckills_count).Count(&count)
	return count
}

func GetSecKillById(id int) (*Seckills, error) {
	var seckill Seckills
	result := db.Where("id = ?", id).Find(&seckill)
	return &seckill, result.Error
}
