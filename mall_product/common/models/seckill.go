package models

func GetSecKillList(pagesize, offsetNum int) (*[]SecKills, error) {
	var seckills []SecKills
	result := db.Limit(pagesize).Offset(offsetNum).Find(&seckills)
	return &seckills, result.Error
}

func GetSecKillCount() (int, error) {
	var count int
	result := db.Find(&[]SecKills{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func AddSeckill(seckill *SecKills) error {
	return db.Table("seckills").Create(seckill).Error
}

func DelSeckill(seckill *SecKills) error {
	return db.Table("seckills").Delete(seckill).Error

}

func GetSeckillInfo(seckill *SecKills) error {
	return db.Table("seckills").First(seckill).Error
}

func UpdateSeckill(id int, kills *SecKills) error {
	return db.Where("id = ?", id).Find(&SecKills{}).Update(kills).Error

}

func GetSeckillByTime(tomorrow_time string, page_size, offset int32) (*[]SecKills, error) {
	var seckills []SecKills
	result := db.Where("start_time <= ?", tomorrow_time).Where("status = ?", 0).Limit(page_size).Offset(offset).Find(&seckills)
	return &seckills, result.Error
}

func GetSecKillCountByTime(tomorrow_time string) int32 {
	seckills_count := []SecKills{}
	var count int32
	db.Where("start_time <= ?", tomorrow_time).Where("status = ?", 0).Find(&seckills_count).Count(&count)
	return count
}

func GetSecKillById(id int) (*SecKills, error) {
	var seckill SecKills
	result := db.Where("id = ?", id).Find(&seckill)
	return &seckill, result.Error
}
