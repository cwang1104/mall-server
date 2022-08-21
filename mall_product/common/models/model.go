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

	SecKills []SecKills
}

type SecKills struct {
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
