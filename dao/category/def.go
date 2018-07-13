package category

// Category category表数据结构
type Category struct {
	ID         uint64 `gorm:"column:id;primary_key:true"`
	Name       string `gorm:"column:name"`
	CreateTime string `gorm:"column:create_time"`
	UpdateTime string `gorm:"column:update_time"`
}

// TableName 修改gorm默认的表名
func (c *Category) TableName() string {
	return "category"
}
