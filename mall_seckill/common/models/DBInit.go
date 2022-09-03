package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"mall_seckill/config"
)

var db *gorm.DB

//初始化数据库连接
func init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)
	//从ini配置文件中加载相应内容
	sec, err := config.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	//tablePrefix = sec.Key("TABLE_PREFIX").String()

	database_link := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err = gorm.Open(dbType, database_link)

	if err != nil {
		log.Println(err)
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return tablePrefix + "_" + defaultTableName
	//}

	db.SingularTable(true) //gorm默认使用复数映射，go代码的单数、复数struct形式都匹配到复数表中：创建表、添加数据时都是如此。指定了db.SingularTable(true)之后，进行严格匹配。
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	fmt.Println("--------数据库初始化成功--------")

}

func CloseDB() {
	defer db.Close()
}
