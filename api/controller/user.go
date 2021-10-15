package controller

import (
	"movie_planet/api/service"
	"movie_planet/models"
	"movie_planet/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user models.UserRegister
	err := c.ShouldBind(&user)
	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Failed binding")
		return
	}

	err = u.service.CreateUser(user)
	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Failed creating user")
		return
	}
	util.SuccessJson(c, http.StatusOK, "Created user")
}

func (u *UserController) LoginUser(c *gin.Context) {
	var user models.UserLogin
	var hmacSecret = []byte(os.Getenv("SECRET_TOKEN"))
	err := c.ShouldBind(&user)
	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Failed")
		return
	}
	dbUser, err := u.service.LoginUser(user)
	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Invalid Login Credentials")
		return
	}
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"user": dbUser,
	// 	"exp":  time.Now().Add(time.Minute * 15).Unix(),
	// })
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.UserClaims{
		User: *dbUser,
	})
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		util.ErrorJson(c, http.StatusBadRequest, "Failed getting token")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Token generated",
		Data:    tokenString,
	}
	c.JSON(http.StatusOK, response)
}

func (u *UserController) UserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": c.MustGet("email").(string),
	})
}
