package initrouter

import (
	_ "lil-helper-backend/docs"
	"lil-helper-backend/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	var Router = gin.Default()

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ApiGroup := Router.Group("")
	router.InitAdminRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)

	return Router
}
