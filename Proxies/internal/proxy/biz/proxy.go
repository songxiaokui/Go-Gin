package biz

/*
@Time    : 2021/3/29 09:35
@Author  : austsxk
@Email   : austsxk@163.com
@File    : proxy.go
@Software: GoLand
*/
import (
	"Go-Gin/Proxies/internal/proxy/data"
	"fmt"
	"gorm.io/gorm"
)

// 这里提供了给server的一系列服务，如登陆 注册 获取代理 添加代理等...
// 因为这里对外提供服务可能需要一些传输对象，比如说新增的实体对象、给前端展示的vo对象都是需要在这里定义

// 新增传输对象
type ProxyDTO struct {
	IP string `json:"ip" form:"ip" binding:"required,ipv4"`
	// 可以为空，但是如果不填，就是0，如果填写，就必须大于4000
	Port    int    `json:"port" form:"port" binding:"omitempty,min=4000,max=8888"`
	Type    string `json:"type" form:"type"`
	Address string `json:"address" form:"address" binding:"min=4,max=8"`
	// 添加一个自定义验证器，使用正则验证这个字段,使用自定义的验证器 ProxyUrlValidate
	ProxyUrl string `json:"url" form:"url" binding:"required,ProxyUrlValidate"`
}

// 查询传输对象
type ProxySearchDTO struct {
	Id    int `json:"id" form:"id" binding:"required"`
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
}

// 返回响应的展示对象VO，有的数据不需要给前端返回，需要进行处理一下
type ProxyVO struct {
}

var _ ProxyServiceInter = (*ProxyServiceImpl)(nil)

// 然后biz层实现具体的方法，给server使用
type ProxyServiceInter interface {
	// 查询代理列表的全部数据,不一定是数据库包含的全部信息
	GetAllProxiesInfo() ([]*data.TbProxy, error)
	// 新增一条数据
	AddOneProxy(*ProxyDTO) error
	// 删除一条数据根据id
	DeleteOneProxyById(id int) error
	// 修改代理数据
	UpdateOneProxy(*ProxyDTO) error
	// 根据id获取一条
	GetProxyById(id int) *data.TbProxy
}

// 定义一个实现,但是要包含底层操作数据库的支持，也就是需要进行接口组合，依赖与data层
type ProxyServiceImpl struct {
	db data.ProxiesDaoInter
}

// 将已经实现ProxiesDaoInter的struct传入，形成一个可操作的对象,将来需要一个实体ProxiesDaoImpl的实现
func MakeProxyService(db data.ProxiesDaoInter) *ProxyServiceImpl {
	return &ProxyServiceImpl{db: db}
}

// 实现上述的业务接口
func (p *ProxyServiceImpl) GetAllProxiesInfo() ([]*data.TbProxy, error) {
	return p.db.GetAllData()
}

func (p *ProxyServiceImpl) AddOneProxy(proxy *ProxyDTO) error {
	// 这里要对参数进行校验的，应该放到服务层,这里假定是已经验证通过
	// 判断重复逻辑还是要在本层处理的
	obj, err := p.db.SearchByIp(proxy.IP)
	fmt.Println(obj)
	if err == nil && obj == nil || err == gorm.ErrRecordNotFound {
		newProxy := &data.TbProxy{
			IP:      proxy.IP,
			Port:    proxy.Port,
			Address: proxy.Address,
			Type:    proxy.Type,
		}
		fmt.Println("新添加的数据:", newProxy)
		err := p.db.Create(newProxy)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return err
}

func (p *ProxyServiceImpl) GetProxyById(id int) *data.TbProxy {
	d, err := p.db.SearchById(id)
	if err != nil {
		return nil
	}
	return d
}

func (p *ProxyServiceImpl) DeleteOneProxyById(id int) error {
	return p.db.DeleteById(id)
}

func (p *ProxyServiceImpl) UpdateOneProxy(*ProxyDTO) error {
	return nil
}
