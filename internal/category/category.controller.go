package category

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

func (a *Controller) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	category, err := Service.Get(id)
	if err != nil {
		fmt.Printf("获取分类失败: %v", err)
		response.FailWithStatus(ctx, http.StatusBadRequest, errors.ErrGetCategory)
		return
	}
	response.RespWithData(ctx, category)

}

func (a *Controller) Create(ctx *gin.Context) {
	var params CreateParamsDto
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("创建分类参数错误: %v", err)
		response.FailWithStatus(ctx, http.StatusBadRequest, errors.ErrCreateCategory)
		return
	}
	if err := Service.Create(ctx, params); err != nil {
		fmt.Printf("创建分类报错: %v", err)
		response.FailWithStatus(ctx, http.StatusBadRequest, errors.ErrCreateCategory)
		return
	}
	response.Ok(ctx)

}
