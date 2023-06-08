package article

import (
	"demo01/pkg/common/errors"
	"demo01/pkg/common/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

var Service = newService()

func (a *Controller) Create(ctx *gin.Context) {
	var params CreateParamsDto
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("创建文章参数错误: %v\n", err)
		response.FailWithStatus(ctx, http.StatusBadRequest, errors.ErrCreateArticle)
		return
	}
	if err := Service.Create(ctx, &params); err != nil {
		fmt.Printf("创建文章失败: %v\n", err)
		response.FailWithStatus(ctx, http.StatusBadRequest, errors.ErrCreateArticle)
		return
	}
	response.Ok(ctx)
}

func (a *Controller) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	article, err := Service.Get(id)
	if err != nil {
		fmt.Printf("获取文章失败: %v\n", err)
		response.Fail(ctx, errors.ErrGetArticle)
		return
	}
	response.RespWithData(ctx, article)
}
