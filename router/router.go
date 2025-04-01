package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"instant_messaging/docs"
	"instant_messaging/middleware"
	"instant_messaging/service"
)

func Router() *gin.Engine {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/index", service.GetIndex)
	r.POST("/login", service.Login)

	r.POST("/user", service.CreateUser)

	r.GET("/sendMsg", service.SendMsg)

	authRoutes := r.Group("/")
	authRoutes.Use(middleware.JWTAuthMiddleware())
	{
		authRoutes.GET("/users", service.GetUserList)
		authRoutes.DELETE("/users/:id", service.DeleteUser)
		authRoutes.PUT("/users/:id", service.UpdateUser)
	}

	return r
}
