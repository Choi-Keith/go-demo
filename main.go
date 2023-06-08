package main

import (
	"demo01/docs"
	"demo01/internal/module"
	"demo01/internal/route"
	"demo01/pkg/gormx"

	"github.com/gin-gonic/gin"
)

//	@title			qinglv
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	moqing.club
//	@contact.email	aksoncai@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:9100
//	@BasePath	/api/v1
func main() {
	engine := gin.New()
	gormx.InitDB(module.Models)

	route.SetupRouter(engine)

	docs.SwaggerInfo.BasePath = "/api/v1"

	engine.Run(":9123")

}
