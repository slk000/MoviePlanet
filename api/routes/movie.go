package routes

import (
	"movie_planet/api/controller"
	"movie_planet/infra"
)

type MovieRoute struct {
	Controller controller.MovieController
	Handler    infra.GinRouter
}

func NewMovieRoute(controller controller.MovieController, handler infra.GinRouter) MovieRoute {
	return MovieRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (m MovieRoute) Setup() {
	movie := m.Handler.Gin.Group("/movies")
	{
		movie.GET("/", m.Controller.GetMovies)
		movie.GET("/:id", m.Controller.GetMovie)
		movie.POST("/", m.Controller.CreateMovie)
		movie.DELETE("/:id", m.Controller.DeleteMovie)
		movie.PUT("/:id", m.Controller.UpdateMovie)

	}
}
