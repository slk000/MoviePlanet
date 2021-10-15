package controller

import (
	"movie_planet/api/service"
	"movie_planet/models"
	"movie_planet/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RateController struct {
	service service.RateService
}

func NewRateController(s service.RateService) RateController {
	return RateController{
		service: s,
	}
}

func (r *RateController) CreateRate(c *gin.Context) {
	var rate models.Rate
	c.ShouldBindJSON(&rate)
	movieID, err := strconv.ParseInt(c.Param("mid"), 10, 64)
	rate.MovieID = movieID
	rate.UserID = 1

	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Failed create rate")
		return
	}
	err = r.service.Create(rate)
	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Failed create rate")
		return
	}
	util.SuccessJson(c, http.StatusCreated, rate)
}

func (r *RateController) GetRates(c *gin.Context) {
	var rates models.Rate
	movieID, err := strconv.ParseInt(c.Param("mid"), 10, 64)
	data, totalRow, err := r.service.FindAll(rates, movieID)

	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Failed find rates")
		return
	}

	responseArray := make([]map[string]interface{}, 0)

	for _, rate := range *data {
		response := rate.ResponseMap()
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
