package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"instant_messaging/docs"
	"instant_messaging/service"
)

func Router() *gin.Engine {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/index", service.GetIndex)
	userGroup := r.Group("/user")
	{
		userGroup.GET("/info", service.GetUserList)
	}

	return r
}
