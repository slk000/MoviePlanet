package main

import (
	"movie_planet/api/controller"
	"movie_planet/api/repo"
	"movie_planet/api/routes"
	"movie_planet/api/service"
	"movie_planet/infra"
	"movie_planet/models"
)

func init() {
	infra.LoadEnv()
}

func main() {
	router := infra.NewGinRouter()
	db := infra.NewDatabase()

	movieRepo := repo.NewMovieRepo(db)
	movieService := service.NewMovieService(movieRepo)
	movieController := controller.NewMovieController(movieService)
	movieRoute := routes.NewMovieRoute(movieController, router)
	movieRoute.Setup()

	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	db.DB.AutoMigrate(&models.Movie{}, &models.User{})
	router.Gin.Run(":80")
}
