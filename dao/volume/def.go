package volume

// Volume volume表结构定义
type Volume struct {
	ID         uint64 `gorm:"column:id;primary_key:true"`
	Name       string `gorm:"column:name"`
	CreateTime string `gorm:"column:create_time"`
	UpdateTime string `gorm:"column:update_time"`
	Status     int    `gorm:"column:status"`
}

// TableName 修改gorm默认的表名
func (v *Volume) TableName() string {
	return "volume"
}
