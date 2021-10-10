package main

import (
	"movie_planet/infra"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		infra.LoadEnv()
		infra.NewDatabase()
		context.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	router.Run(":80")
}
