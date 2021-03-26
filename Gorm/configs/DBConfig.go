package configs

/*
@Time    : 2021/3/26 20:57
@Author  : austsxk
@Email   : austsxk@163.com
@File    : DBConfig.go
@Software: GoLand
*/

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	// 定义连接数据库url
	DbUrl = "root:mysql@tcp(127.0.0.1:3306)/proxies?charset=utf8mb4&parseTime=True&loc=Local"
)

var SqlDb *gorm.DB

func InitDb() error {
	db, err := gorm.Open(mysql.Open(DbUrl), &gorm.Config{})
	if err != nil {
		return err
	}
	SqlDb = db
	return nil
}
