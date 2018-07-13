package content

import (
	"github.com/FrontMage/HelloGithubNavBackend/dao"
	"github.com/FrontMage/HelloGithubNavBackend/dao/category"
	"github.com/FrontMage/HelloGithubNavBackend/dao/volume"
)

// Content content表的数据定义
type Content struct {
	ID          uint64            `gorm:"column:id,primary_key:true"`
	ProjectURL  string            `gorm:"column:project_url"`
	Title       string            `gorm:"column:title"`
	Description string            `gorm:"description"`
	ImagePath   string            `gorm:"image_path"`
	CategoryID  uint64            `gorm:"category_id"`
	VolumeID    uint64            `gorm:"volume_id"`
	CreateTime  string            `gorm:"column:create_time"`
	UpdateTime  string            `gorm:"column:update_time"`
	Status      int               `gorm:"column:status"`
	Category    category.Category `gorm:"foreignkey:ID;association_foreignkey:CategoryID"`
	Volume      volume.Volume     `gorm:"foreignkey:ID;association_foreignkey:VolumeID"`
}

// TableName 修改gorm默认的表名
func (c *Content) TableName() string {
	return "content"
}

// Get 根据主键id获取一条记录
func Get(id uint64) (*Content, error) {
	c := &Content{ID: id}
	if err := dao.DB.
		// Debug().
		Where("id = ?", id).
		Preload("Category").
		Preload("Volume").
		First(&c).
		Error; err != nil {
		return nil, err
	}
	return c, nil
}
