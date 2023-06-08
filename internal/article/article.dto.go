package article

import "demo01/internal/category"

type CreateParamsDto struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Content     string             `json:"content"`
	CategoryID  string             `json:"categoryID"`
	Category    *category.Category `json:"category"`
}
