package models

import "time"

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	Num         int       `json:"num"`
	Unit        string    `json:"unit"`
	Picture     string    `json:"picture"`
	Desc        string    `json:"desc"`
	CreatedTime time.Time `json:"created_time"`

	SecKills []Seckills
}

type Seckills struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	Num       int       `json:"num"`
	GoodsId   int       `json:"goods_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	// 1表示下架，0表示未下架
	Status      int       `json:"status"`
	CreatedTime time.Time `json:"created_time"`
}

type Order struct {
	ID          int       `json:"id"`
	OrderNum    int       `json:"order_num"`
	UserId      int       `json:"user_id"`
	ActivityId  int       `json:"activity_id"`
	PayStatus   int       `json:"pay_status"`
	CreatedTime time.Time `json:"created_time"`
}
