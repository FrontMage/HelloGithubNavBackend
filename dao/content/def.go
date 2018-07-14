package content

import (
	"github.com/FrontMage/HelloGithubNavBackend/dao"
	"github.com/FrontMage/HelloGithubNavBackend/dao/category"
	"github.com/FrontMage/HelloGithubNavBackend/dao/volume"
)

// Content content表的数据定义
type Content struct {
	ID          uint64             `gorm:"column:id,primary_key:true" json:"id,omitempty"`
	ProjectURL  string             `gorm:"column:project_url" json:"projectURL,omitempty"`
	Title       string             `gorm:"column:title" json:"title,omitempty"`
	Description string             `gorm:"description" json:"description,omitempty"`
	ImagePath   string             `gorm:"image_path" json:"imagePath,omitempty"`
	CategoryID  uint64             `gorm:"category_id" json:"categoryID,omitempty"`
	VolumeID    uint64             `gorm:"volume_id" json:"volumeID,omitempty"`
	CreateTime  string             `gorm:"column:create_time" json:"createTime,omitempty"`
	UpdateTime  string             `gorm:"column:update_time" json:"updateTime,omitempty"`
	Status      int                `gorm:"column:status" json:"status,omitempty"`
	Category    *category.Category `gorm:"foreignkey:ID;association_foreignkey:CategoryID" json:"category,omitempty"`
	Volume      *volume.Volume     `gorm:"foreignkey:ID;association_foreignkey:VolumeID" json:"volume,omitempty"`
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

// BatchGet 批量获取记录
func BatchGet(ids []uint64) ([]*Content, error) {
	c := []*Content{}
	return c, dao.DB.
		// Debug().
		Preload("Category").
		Preload("Volume").
		Where("id in (?)", ids).
		Find(&c).Error
}
