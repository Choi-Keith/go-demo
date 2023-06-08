package category

import "demo01/internal/base"

type Category struct {
	base.Model
	ID          string `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:128;index:unique" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	CreatorID   string `gorm:"size:128" json:"creatorID"`
	CreatorName string `gorm:"size:128" json:"creatorName"`
}
