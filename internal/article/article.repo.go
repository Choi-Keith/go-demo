package article

import "demo01/pkg/gormx"

type repo struct {
}

func newRepo() *repo {
	return &repo{}
}

func (a *repo) GetArticleByTitle(title string) (*Article, error) {
	var article Article
	if err := gormx.DB.Take(&article, "title=?", title).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (a *repo) Create(params *Article) error {
	return gormx.DB.Create(params).Error
}

func (a *repo) Get(id string) (*Article, error) {
	article := &Article{}
	if err := gormx.DB.Preload("Category").First(article, "articles.id=?", id).Error; err != nil {
		return nil, err
	}
	return article, nil
}
