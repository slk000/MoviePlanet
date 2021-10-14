package routes

import (
	"movie_planet/api/controller"
	"movie_planet/infra"
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
	user := u.Handler.Gin.Group("/auth")
	{
		user.POST("/register", u.Controller.CreateUser)
		user.POST("/login", u.Controller.LoginUser)
	}
}
