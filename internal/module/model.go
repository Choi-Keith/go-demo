package module

import (
	"demo01/internal/article"
	"demo01/internal/category"
	"demo01/internal/user"
)

var Models = []interface{}{
	&user.User{},
	&category.Category{},
	&article.Article{},
}
