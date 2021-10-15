package routes

import (
	"movie_planet/api/controller"
	"movie_planet/infra"
	"movie_planet/middlewares"
)

type UserRoute struct {
	Handler    infra.GinRouter
	Controller controller.UserController
}

func NewUserRoute(controller controller.UserController, handler infra.GinRouter) UserRoute {
	return UserRoute{
		Handler:    handler,
		Controller: controller,
	}
}

func (u UserRoute) Setup() {
	userAuth := u.Handler.Gin.Group("/auth")
	{
		userAuth.POST("/register", u.Controller.CreateUser)
		userAuth.POST("/login", u.Controller.LoginUser)
	}

	user := u.Handler.Gin.Group("/user")
	{
		user.GET("/", middlewares.AuthMiddleware(), u.Controller.UserInfo)
	}
}
