package article

import (
	category2 "demo01/internal/category"
	"demo01/util"
	"errors"

	"github.com/gin-gonic/gin"
)

type service struct {
}

func newService() *service {
	return &service{}
}

var Repo = newRepo()

func (a *service) getArticleByTitle(title string) (*Article, error) {
	return Repo.GetArticleByTitle(title)
}

func (a *service) Create(ctx *gin.Context, params *CreateParamsDto) error {
	if params.Title == "" {
		return errors.New("文章标题不为空")
	}
	ab, _ := a.getArticleByTitle(params.Title)
	if ab != nil {
		return errors.New("文章标题已存在")
	}

	if params.CategoryID == "" {
		return errors.New("文章分类不为空")
	}

	category, _ := category2.Repo.Get(params.CategoryID)
	if category == nil {
		return errors.New("文章分类不存在")
	}

	creatorID, _ := ctx.Get("userID")
	creatorName, _ := ctx.Get("username")
	article := &Article{
		ID:          util.GenUUID(""),
		Title:       params.Title,
		Description: params.Description,
		Content:     params.Content,
		CategoryID:  params.CategoryID,
		CreatorID:   creatorID.(string),
		CreatorName: creatorName.(string),
	}

	err := Repo.Create(article)
	return err
}

func (a *service) Get(id string) (*Article, error) {
	return Repo.Get(id)
}
