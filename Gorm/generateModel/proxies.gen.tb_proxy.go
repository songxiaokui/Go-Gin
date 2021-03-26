package generateModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TbProxyMgr struct {
	*_BaseMgr
}

// TbProxyMgr open func
func TbProxyMgr(db *gorm.DB) *_TbProxyMgr {
	if db == nil {
		panic(fmt.Errorf("TbProxyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TbProxyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tb_proxy"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TbProxyMgr) GetTableName() string {
	return "tb_proxy"
}

// Get 获取
func (obj *_TbProxyMgr) Get() (result TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TbProxyMgr) Gets() (results []*TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TbProxyMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithIP ip获取
func (obj *_TbProxyMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithPort port获取
func (obj *_TbProxyMgr) WithPort(port int) Option {
	return optionFunc(func(o *options) { o.query["port"] = port })
}

// WithType type获取
func (obj *_TbProxyMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithAddress address获取
func (obj *_TbProxyMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// GetByOption 功能选项模式获取
func (obj *_TbProxyMgr) GetByOption(opts ...Option) (result TbProxy, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TbProxyMgr) GetByOptions(opts ...Option) (results []*TbProxy, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_TbProxyMgr) GetFromID(id int) (result TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_TbProxyMgr) GetBatchFromID(ids []int) (results []*TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容
func (obj *_TbProxyMgr) GetFromIP(ip string) (results []*TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量唯一主键查找
func (obj *_TbProxyMgr) GetBatchFromIP(ips []string) (results []*TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip IN (?)", ips).Find(&results).Error

	return
}

// GetFromPort 通过port获取内容
func (obj *_TbProxyMgr) GetFromPort(port int) (results []*TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("port = ?", port).Find(&results).Error

	return
}

// GetBatchFromPort 批量唯一主键查找
func (obj *_TbProxyMgr) GetBatchFromPort(ports []int) (results []*TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("port IN (?)", ports).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_TbProxyMgr) GetFromType(_type string) (results []*TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找
func (obj *_TbProxyMgr) GetBatchFromType(_types []string) (results []*TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容
func (obj *_TbProxyMgr) GetFromAddress(address string) (results []*TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量唯一主键查找
func (obj *_TbProxyMgr) GetBatchFromAddress(addresss []string) (results []*TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address IN (?)", addresss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TbProxyMgr) FetchByPrimaryKey(id int) (result TbProxy, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
