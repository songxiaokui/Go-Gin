package biz

import (
	"Go-Gin/Gorm/configs"
	"gorm.io/gorm"
)

/*
@Time    : 2021/3/26 21:57
@Author  : austsxk
@Email   : austsxk@163.com
@File    : Users.go
@Software: GoLand
*/

type Proxies struct {
	// embed
	*BaseModel
	ID      int    `gorm:"column:id;primaryKey;autoIncrement"`
	IP      string `gorm:"column:ip"`
	Port    int    `gorm:"column:port"`
	Type    string `gorm:"column:type"`
	Address string `gorm:"column:address"`
}

// 实现 Tabler 接口，可以直接实现返回表名字
func (*Proxies) TableName() string {
	return "tb_proxy"
}

// 构造对象
func NewProxies() *Proxies {
	p := &Proxies{}
	p.BaseModel = Base(p)
	return p
}

// 根据Port获取
func (p *Proxies) LoadByPort(port int) *Proxies {
	p.Port = port
	configs.SqlDb.Where(p).First(p)
	return p
}

// 暂时只针对一个字段进行的编写，后面会抽像一个基类，进行封装
func (p *Proxies) IdCompareGenerate(value int, opt int) CompareFunc {
	return func(db *gorm.DB) *gorm.DB {
		switch opt {
		case CompareGraterThan:
			return db.Where("id > ?", value)
		case CompareLessThan:
			return db.Where("id < ?", value)
		case CompareGraterEqual:
			return db.Where("id >= ?", value)
		case CompareLessEqual:
			return db.Where("id <= ?", value)
		case ComPareLike:
			return db.Where("id like ?", value)
		default:
			return db.Where("id = ?", value)
		}
	}
}

func (p *Proxies) IdCompare(value int, opt int) CompareFunc {
	return p.IdCompareGenerate(value, opt)
}

func (p *Proxies) Filter(cf ...func(*gorm.DB) *gorm.DB) {
	configs.SqlDb.Scopes(cf...).First(p)
}
