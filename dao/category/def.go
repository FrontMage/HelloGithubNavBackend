package category

import "github.com/FrontMage/HelloGithubNavBackend/dao"

// Category category表数据结构
type Category struct {
	ID         uint64 `gorm:"column:id;primary_key:true" json:"id,omitempty"`
	Name       string `gorm:"column:name" json:"name,omitempty"`
	CreateTime string `gorm:"column:create_time" json:"createTime,omitempty"`
	UpdateTime string `gorm:"column:update_time" json:"updateTime,omitempty"`
}

// TableName 修改gorm默认的表名
func (c *Category) TableName() string {
	return "category"
}

// Get get by id
func Get(id uint64) (*Category, error) {
	c := &Category{}
	return c, dao.DB.Where("id = ?", id).First(c).Error
}
