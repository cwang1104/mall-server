package models

import "time"

type Admin struct {
	Id          int       `json:"id"`
	UserName    string    `json:"user_name"`
	Password    string    `json:"password"`
	Desc        string    `json:"desc"`
	Status      int       `json:"status"`
	CreatedTime time.Time `json:"created_time"`
}
