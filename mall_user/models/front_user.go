package models

import (
	"fmt"
	"time"
)

type User struct {
	Id          int       `json:"id" gorm:"primary_key"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"password"`
	Desc        string    `json:"desc"`
	Status      int       `json:"status"`
	CreatedTime time.Time `json:"created_time"`
}

func RegisterUser(user *User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func EmaiIsExist(email string) bool {
	var count int
	db.Where("email = ?", email).Find(&User{}).Count(&count)
	if count < 1 {
		return false
	} else {
		return true
	}

}

func FindUserByEmail(email string) *User {
	var user User
	db.Where("email = ?", email).First(&user)
	return &user
}

func GetUserLists(pageSize, offsetNum int) (*[]User, error) {
	var users []User
	result := db.Limit(pageSize).Offset(offsetNum).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func GetUserCount() *int {
	var count int
	db.Find(&User{}).Count(&count)
	return &count
}

func AddUserInfo() {
	a := 12345
	for i := 0; i < 100; i++ {
		a = a + 1
		user := User{
			Email:       fmt.Sprintf("%d@qq.com", a),
			Password:    "c0264d69d070404c22585b842fb642bc",
			Desc:        "测试用户",
			Status:      1,
			CreatedTime: time.Now(),
		}
		db.Create(&user)
	}
}
