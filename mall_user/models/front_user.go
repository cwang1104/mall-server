package models

import "time"

type User struct {
	Id          int       `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Desc        string    `json:"desc"`
	Status      int       `json:"status"`
	CreatedTime time.Time `json:"created_time"`
}
