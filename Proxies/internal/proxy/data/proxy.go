package data

import (
	"Go-Gin/Proxies/configs"
	"fmt"
)

/*
@Time    : 2021/3/29 10:27
@Author  : austsxk
@Email   : austsxk@163.com
@File    : proxy.go
@Software: GoLand
*/

type TbProxy struct {
	ID      int    `gorm:"primaryKey;column:id;type:int(11);not null;AUTO_INCREMENT" json:"-"`
	IP      string `gorm:"column:ip;type:varchar(255)" json:"ip"`
	Port    int    `gorm:"column:port;type:int(11)" json:"port"`
	Type    string `gorm:"column:type;type:varchar(255)" json:"type"`
	Address string `gorm:"column:address;type:varchar(255)" json:"address"`
}

func (*TbProxy) TableName() string {
	return "tb_proxy"
}

// 验证接口是否实现
var _ ProxiesDaoInter = (*ProxiesDaoImpl)(nil)

// 定义了需要给业务提供的接口
type ProxiesDaoInter interface {
	// 添加一条新的代理信息保存在数据库中
	Create(proxy *TbProxy) error
	// 根据id查询数据
	SearchById(id int) (*TbProxy, error)
	// 获取全部信息
	GetAllData() ([]*TbProxy, error)
	// 根据ip查询是否存在数据
	SearchByIp(ip string) (*TbProxy, error)
	// 删除一条数据
	DeleteById(id int) error
}

// 定义了一个接口，就要有一个实现对象
type ProxiesDaoImpl struct {
}

func (p *ProxiesDaoImpl) Create(proxy *TbProxy) error {
	fmt.Println("data层接受的数据:", proxy)
	return configs.SqlDb.Create(&proxy).Error
}

func (p *ProxiesDaoImpl) SearchById(id int) (*TbProxy, error) {
	proxy := &TbProxy{}
	proxy.ID = id
	err := configs.SqlDb.Find(&proxy).Error
	if err != nil {
		return nil, err
	}
	return proxy, nil
}

func (p *ProxiesDaoImpl) GetAllData() ([]*TbProxy, error) {
	var proxy []*TbProxy
	err := configs.SqlDb.Find(&proxy).Error
	if err != nil {
		return nil, err
	}
	return proxy, nil
}

func (p *ProxiesDaoImpl) SearchByIp(ip string) (*TbProxy, error) {
	proxy := &TbProxy{}
	err := configs.SqlDb.Where("ip = ?", ip).First(proxy).Error
	if err != nil {
		return nil, err
	}
	return proxy, nil
}

func (p *ProxiesDaoImpl) DeleteById(id int) error {
	proxy := &TbProxy{}
	proxy.ID = id
	return configs.SqlDb.Delete(&proxy).Error
}
