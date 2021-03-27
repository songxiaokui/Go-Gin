package biz

import (
	"fmt"
	"gorm.io/gorm"
)

/*
@Time    : 2021/3/27 15:00
@Author  : austsxk
@Email   : austsxk@163.com
@File    : compareUtility.go
@Software: GoLand
*/

// 抽象通用的方法
type CompareFunc func(db *gorm.DB) *gorm.DB

const (
	CompareGraterThan  = 10 // 大于
	CompareLessThan    = 11 // 小于
	CompareEqual       = 12 // 等于
	CompareGraterEqual = 13 // 大于等于
	CompareLessEqual   = 14 // 小于等于
	ComPareLike        = 15 // 模糊查询
)

// 根据ID进行关系查询 大于 小于 等于, 并且传入指定的字段，将其封装
func generateCompareFunc(filed string, value interface{}, opt int) CompareFunc {
	return func(db *gorm.DB) *gorm.DB {
		switch opt {
		case CompareGraterThan:
			return db.Where(fmt.Sprintf("%s > ?", filed), value)
		case CompareLessThan:
			return db.Where(fmt.Sprintf("%s < ?", filed), value)
		case CompareGraterEqual:
			return db.Where(fmt.Sprintf("%s >= ?", filed), value)
		case CompareLessEqual:
			return db.Where(fmt.Sprintf("%s <= ?", filed), value)
		case ComPareLike:
			return db.Where(fmt.Sprintf("%s like ?", filed), value)
		default:
			return db.Where(fmt.Sprintf("%s = ?", filed), value)
		}
	}
}
