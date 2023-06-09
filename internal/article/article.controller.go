package article

import (
	"context"
	"demo01/pkg/common/errors"
	"demo01/pkg/common/response"
	"demo01/pkg/logger/logrusx"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

var Service = newService()

func (a *Controller) Create(ctx *gin.Context) {
	var params CreateParamsDto
	if err := ctx.ShouldBindJSON(&params); err != nil {
		logrusx.Errorf(context.Background(), "创建文章参数错误: %v\n", err)
		response.FailWithStatus(ctx, http.StatusBadRequest, errors.ErrCreateArticle)
		return
	}
	if err := Service.Create(ctx, &params); err != nil {
		logrusx.Errorf(context.Background(), "创建文章失败: %v\n", err)
		response.FailWithStatus(ctx, http.StatusBadRequest, errors.ErrCreateArticle)
		return
	}
	response.Ok(ctx)
}

func (a *Controller) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	article, err := Service.Get(id)
	if err != nil {
		logrusx.Errorf(ctx, "获取文章失败: %v\n", err)
		response.Fail(ctx, errors.ErrGetArticle)
		return
	}
	response.RespWithData(ctx, article)
}

func (a *Controller) Insert(ctx *gin.Context) {
	Service.Insert()
	response.Ok(ctx)
}
