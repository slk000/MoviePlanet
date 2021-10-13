package controller

import (
	"movie_planet/api/service"
	"movie_planet/models"
	"movie_planet/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	service service.MovieService
}

func NewMovieController(s service.MovieService) MovieController {
	return MovieController{
		service: s,
	}
}

func (m MovieController) GetMovies(c *gin.Context) {
	var movies models.Movie
	keyword := c.Query("keyword")
	data, totalRow, err := m.service.FindAll(movies, keyword)

	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Failed find movies")
		return
	}

	responseArray := make([]map[string]interface{}, 0)

	for _, n := range *data {
		response := n.ResponseMap()
		responseArray = append(responseArray, response)
	}

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "OK",
		Data: map[string]interface{}{
			"rows":       responseArray,
			"total_rows": totalRow,
		},
	})
}

func (m *MovieController) GetMovie(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var movie models.Movie
	movie.ID = id
	res, err := m.service.Find(movie)
	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Failed finding movie")
		return
	}
	response := res.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "OK",
		Data:    &response,
	})
}

func (m *MovieController) CreateMovie(c *gin.Context) {
	var movie models.Movie
	c.ShouldBindJSON(&movie)

	if movie.Name == "" {
		util.ErrorJson(c, http.StatusBadRequest, "Name is required")
		return
	}
	err := m.service.Create(movie)
	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Failed to create movie")
		return
	}
	util.SuccessJson(c, http.StatusCreated, "Created")
}

func (m *MovieController) DeleteMovie(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	err = m.service.Delete(id)

	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Failed deleting")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted",
	}
	c.JSON(http.StatusOK, response)
}

func (m *MovieController) UpdateMovie(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	var movie models.Movie
	movie.ID = id

	oldMovie, err := m.service.Find(movie)

	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Movie not found")
		return
	}
	c.ShouldBindJSON(&oldMovie)
	// todo

	util.ErrorJson(c, http.StatusNotImplemented, "Not implemented")
}
