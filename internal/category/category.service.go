package category

import (
	"demo01/util"
	"errors"

	"github.com/gin-gonic/gin"
)

type service struct {
}

var Repo = newRepo()

func newService() *service {
	return &service{}
}

func (a *service) getCategoryByName(name string) (*Category, error) {
	return a.getCategoryByName(name)
}

func (a *service) Create(ctx *gin.Context, params CreateParamsDto) error {
	if len(params.Name) == 0 {
		return errors.New("请输入标签名")
	}

	c, _ := a.getCategoryByName(params.Name)
	if c != nil {
		return errors.New("名称已存在")
	}
	id := util.GenUUID("")
	creatorID, ok := ctx.Get("userID")
	if !ok {
		return errors.New("获取创建者ID失败")
	}
	creatorName, ok := ctx.Get("username")
	if !ok {
		return errors.New("获取创建者名称失败")
	}
	category := &Category{
		ID:          id,
		Name:        params.Name,
		Description: params.Description,
		CreatorID:   creatorID.(string),
		CreatorName: creatorName.(string),
	}
	return Repo.Create(category)
}

func (a *service) Get(id string) (*Category, error) {
	return Repo.Get(id)
}
