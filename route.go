package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-bai/forward/controller"
	"github.com/go-bai/forward/middleware"
)

func setRoute(route *gin.Engine) {
	api := route.Group("api")
	{
		api.POST("/users", controller.Register)
		api.POST("/users/login", controller.Login)
		authAPI := api.Use(middleware.Auth())
		{
			authAPI.GET("/users", controller.UserList)
			authAPI.GET("/user", controller.UserDetail)
		}
	}
}
