package route

import (
	"demo01/internal/article"
	"demo01/internal/category"
	"demo01/internal/user"
	"demo01/middleware"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {

	engine.Use(middleware.CORSMiddleware(), gin.Recovery())
	v1 := engine.Group("/api/v1")

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	oauthGroup := v1.Group("oauth")
	oauthController := &user.Controller{}
	{
		oauthGroup.GET("/:provider/authorize", oauthController.Authorize)
		oauthGroup.GET("/:provider/callback", oauthController.LoginOauth)
	}

	authGroup := v1.Group("auth")
	authController := &user.Controller{}
	{
		authGroup.POST("/register", authController.Create)
		authGroup.POST("/login", authController.Login)
	}

	jwtAuth := v1.Group("")
	jwtAuth.Use(middleware.JWTMiddleware())

	userController := &user.Controller{}
	userGroup := jwtAuth.Group("user")
	{
		userGroup.GET("", userController.List)
		userGroup.GET("/:id", userController.Get)
	}

	categoryController := &category.Controller{}
	categoryGroup := jwtAuth.Group("category")
	{
		categoryGroup.GET("/:id", categoryController.Get)
		categoryGroup.POST("", categoryController.Create)
	}

	articleController := &article.Controller{}
	articleGroup := jwtAuth.Group("article")
	{
		articleGroup.GET("/:id", articleController.Get)
		articleGroup.POST("", articleController.Create)
	}

}
