package models

import "time"

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
