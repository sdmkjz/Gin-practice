package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identitu;type:varchar(36);" json:"identity"`       // 问题表的标识
	CategoryId string `gorm:"column:category_id;type:varchar(255)" json:"category_id"` // 分类ID
	Title      string `gorm:"column:title;type:varchar(255);" json:"title"`            // 文章标题
	Content    string `gorm:"column:content;type:text;" json:"content"`                // 文章正文
}

func (table Problem) TableName() string {
	return ""
}
