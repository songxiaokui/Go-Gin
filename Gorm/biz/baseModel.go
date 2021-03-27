package biz

import (
	"Go-Gin/Gorm/configs"
	"gorm.io/gorm"
)

/*
@Time    : 2021/3/27 15:20
@Author  : austsxk
@Email   : austsxk@163.com
@File    : baseModel.go
@Software: GoLand
*/

// 抽象一个父类，封装父亲方法
type BaseModel struct {
	child interface{}
}

func Base(v interface{}) *BaseModel {
	return &BaseModel{
		child: v,
	}
}

// 将子类的方法抽象到父类
// 根据ID获取数据
func (b *BaseModel) LoadById(id int) interface{} {
	configs.SqlDb.First(b.child, id)
	return b.child
}

// 比较查询
func (b *BaseModel) CompareByFiled(filed string, value interface{}, opt int) CompareFunc {
	return generateCompareFunc(filed, value, opt)
}

func (b *BaseModel) Filter(cf ...func(*gorm.DB) *gorm.DB) {
	configs.SqlDb.Scopes(cf...).First(b.child)
}
