package response

import (
	"demo01/pkg/common/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": errors.ActionSuccess.Code,
		"msg":  errors.ActionSuccess.Message,
		"data": struct{}{},
	})
}

func RespWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": errors.ActionSuccess.Code,
		"msg":  errors.ActionSuccess.Message,
		"data": data,
	})
}

func RespResourceNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": errors.ErrorNotFound.Code,
		"msg":  errors.ErrorNotFound.Message,
		"data": struct{}{},
	})
}

func RespWithHttpCode(c *gin.Context, httpCode int, err *errors.CodeMessage) {
	c.JSON(httpCode, gin.H{
		"code": err.Code,
		"msg":  err.Message,
	})
}

func Fail(c *gin.Context, err *errors.CodeMessage) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": err.Code,
		"msg":  err.Message,
		"data": struct{}{},
	})
}

func FailWithStatus(c *gin.Context, status int, err *errors.CodeMessage) {
	c.JSON(status, gin.H{
		"code": err.Code,
		"msg":  err.Message,
		"data": struct{}{},
	})
}
