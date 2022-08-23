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
