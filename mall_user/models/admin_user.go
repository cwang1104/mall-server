package models

import (
	"fmt"
	"time"
)

type Admin struct {
	Id          int       `json:"id"`
	UserName    string    `json:"user_name"`
	Password    string    `json:"password"`
	Desc        string    `json:"desc"`
	Status      int       `json:"status"`
	CreatedTime time.Time `json:"created_time"`
}

func IsAdminExist(user_name string) bool {
	var count int
	db.Where("user_name = ?", user_name).Find(&Admin{}).Count(&count)
	if count < 1 {
		return false
	} else {
		return true
	}

}

func FindAdminByUserName(user_name string) *Admin {
	var admin Admin
	db.Where("user_name = ?", user_name).First(&admin)
	fmt.Println(admin)
	return &admin
}
