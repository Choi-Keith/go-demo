package article

import (
	"demo01/internal/base"
	"demo01/internal/category"
)

type Article struct {
	base.Model
	ID          string             `gorm:"primaryKey" json:"id"`
	Title       string             `gorm:"size:128;index:unique" json:"title"`
	Description string             `gorm:"type:text" json:"description"`
	Content     string             `gorm:"type:text" json:"content"`
	CreatorID   string             `gorm:"size:128;index:idx_creator_id" json:"creatorID"`
	CreatorName string             `gorm:"size:128" json:"creatorName"`
	CategoryID  string             `gorm:"size:128;index:idx_category_id" json:"categoryID"`
	Category    *category.Category `json:"category"`
}
