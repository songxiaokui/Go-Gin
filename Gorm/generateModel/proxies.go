package generateModel

// TbProxy [...]
type TbProxy struct {
	ID      int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	IP      string `gorm:"column:ip;type:varchar(255)" json:"ip"`
	Port    int    `gorm:"column:port;type:int(11)" json:"port"`
	Type    string `gorm:"column:type;type:varchar(255)" json:"type"`
	Address string `gorm:"column:address;type:varchar(255)" json:"address"`
}
