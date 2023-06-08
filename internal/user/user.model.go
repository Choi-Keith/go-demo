package user

import (
	"database/sql"
	"demo01/internal/base"
)

type User struct {
	base.Model
	ID          string         `gorm:"primaryKey" json:"id"`
	Username    sql.NullString `gorm:"size:128;index:unique" json:"username" form:"username"`
	Email       sql.NullString `gorm:"size:128;index:unique" json:"email" form:"email"`
	Phone       sql.NullString `gorm:"size:32;index:unique" json:"phone" form:"phone"`
	WeChatID    sql.NullString `gorm:"size:128;index:unique" json:"weChatID" form:"weChatID"`
	Password    string         `gorm:"size:512" json:"-" form:"-"`
	Website     string         `gorm:"size:1024" json:"website" form:"website"`
	Avatar      string         `gorm:"type:text" json:"avatar" form:"avatar"`
	Description string         `gorm:"type:text" json:"description" form:"description"`
	Level       int            `gorm:"not null;default 0" json:"level" form:"level"`
	CreatorName string         `gorm:"size:128;index:unique" json:"creatorName"`
	CreatorID   string         `json:"creatorID"`
}
