package biz

import "Go-Gin/Gorm/configs"

/*
@Time    : 2021/3/26 21:57
@Author  : austsxk
@Email   : austsxk@163.com
@File    : Users.go
@Software: GoLand
*/

type Proxies struct {
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
	return &Proxies{}
}

// 根据ID获取数据
func (p *Proxies) LoadById(id int) *Proxies {
	configs.SqlDb.First(p, id)
	return p
}

// 根据Port获取
func (p *Proxies) LoadByPort(port int) *Proxies {
	p.Port = port
	configs.SqlDb.Where(p).First(p)
	return p
}
