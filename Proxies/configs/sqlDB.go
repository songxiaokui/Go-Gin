package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

/*
@Time    : 2021/3/29 10:22
@Author  : austsxk
@Email   : austsxk@163.com
@File    : sqlDB.go
@Software: GoLand
*/

// 数据库的连接
const DBUrl = "root:mysql@tcp(127.0.0.1:3306)/proxies?charset=utf8mb4&parseTime=True&loc=Local"

var SqlDb *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open(DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	SqlDb = db
}
