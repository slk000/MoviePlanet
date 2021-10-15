package routes

import (
	"movie_planet/api/controller"
	"movie_planet/infra"
)

type RateRoute struct {
	Controller controller.RateController
	Handler    infra.GinRouter
}

func NewRateRoute(controller controller.RateController, handler infra.GinRouter) RateRoute {
	return RateRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (r RateRoute) Setup(movieRoute *MovieRoute) {
	// rate := movieRoute.Handler.Gin.RouterGroup.Group("/:mid/rates")
	rate := r.Handler.Gin.Group("/rates/:mid/")
	{
		rate.GET("/", r.Controller.GetRates)
		// rate.GET("/:rid", r.Controller.GetRate)
		rate.POST("/", r.Controller.CreateRate)
	}
}
