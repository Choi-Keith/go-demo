package base

type Model struct {
	CreatedAt int64 `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"updatedAt"`
	DeletedAt int64 `json:"-"`
}
