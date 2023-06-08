package user

import (
	"demo01/pkg/common/errors"
	"demo01/pkg/common/response"
	"demo01/pkg/oauth/gitee"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *service
}

func (a *Controller) Login(ctx *gin.Context) {
	var params LoginParamsDto
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		fmt.Println("用户或密码错误", err)
		response.FailWithStatus(ctx, http.StatusBadRequest, errors.ErrorParamsLogin)
		return
	}
	loginRes := &LoginRespDto{}
	loginRes, err = a.Service.Login(params)
	if err != nil {
		fmt.Println("用户不存在", err)
		response.FailWithStatus(ctx, http.StatusBadRequest, errors.ErrorNotExistsLogin)
		return
	}
	response.RespWithData(ctx, loginRes)

}

func (a *Controller) Authorize(ctx *gin.Context) {
	//provider := ctx.Param("provider")

	url := gitee.AuthCodeURL(map[string]string{})
	response.RespWithData(ctx, url)
}

func (a *Controller) LoginOauth(ctx *gin.Context) {
	//provider := ctx.Param("provider")

	var oauthDto LoginOAuthParamsDto

	err := ctx.ShouldBindQuery(&oauthDto)
	if err != nil {
		fmt.Println("loginOauth============111", err)
		response.Fail(ctx, errors.ErrLoginOauthParams)
		return
	}
	loginRes, err := Service.LoginOAuth(oauthDto.Code, oauthDto.State)
	if err != nil {
		fmt.Println("loginOauth============222", err)
		response.Fail(ctx, errors.ErrLoginOauthParams)
		return
	}
	response.RespWithData(ctx, loginRes)
}

// List godoc
// @Summary 列表
// @Schemes
// @Description 列表
// @Tags 用户
// @Accept json
// @Produce json
// @Param params query ListParamsDto true	"list search by params"
// @Success 200 {object} ListRespDto
// @Router /user [get]
func (a *Controller) List(ctx *gin.Context) {
	var params ListParamsDto
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		response.Fail(ctx, &errors.CodeMessage{
			Code:    400,
			Message: err.Error(),
		})
	}
	response.RespWithData(ctx, map[string]interface{}{
		"total": 10,
	})
}

func (a *Controller) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := a.Service.Get(id)
	if err != nil {
		response.Fail(ctx, errors.ErrorNotFound)
		return
	}
	response.RespWithData(ctx, user)
}

// Create godoc
// @Summary 创建用户
// @Schemes
// @Description 创建用户
// @Param create body CreateParamsDto true "Create"
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} CreateRespDto
// @Router /user [post]
func (a *Controller) Create(ctx *gin.Context) {
	var user CreateParamsDto
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	_, err = a.Service.Create(ctx, user)
	if err != nil {
		fmt.Println("创建用户失败: ", err)
		response.FailWithStatus(ctx, http.StatusBadRequest, errors.ErrorCreateUser)
		return
	}
	response.Ok(ctx)
}
