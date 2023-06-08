package middleware

import (
	"demo01/pkg/common/errors"
	"demo01/pkg/common/response"
	"demo01/pkg/jwtx"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := jwtx.GetToken(ctx)
		claims, err := jwtx.ParseToken(tokenString)
		if err != nil {
			fmt.Printf("ParseToken failed: %v", err)
			response.FailWithStatus(ctx, http.StatusUnauthorized, errors.ErrorMissingJwtToken)
			ctx.Abort()
			return
		}
		sub := strings.Split(claims.Subject, "-")
		ctx.Set("userID", sub[0])
		ctx.Set("username", sub[1])
		ctx.Next()
	}
}
