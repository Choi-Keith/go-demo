package category

import "demo01/pkg/gormx"

type repo struct {
}

func newRepo() *repo {
	return &repo{}
}

func (a *repo) GetCategoryByName(name string) (*Category, error) {
	var category *Category
	if err := gormx.DB.Take(&category, "name=?", name).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (a *repo) Get(id string) (*Category, error) {
	category := &Category{}
	err := gormx.DB.First(category, "where id=?", id).Error
	return category, err
}

func (a *repo) Create(params *Category) error {
	return gormx.DB.Create(params).Error
}
