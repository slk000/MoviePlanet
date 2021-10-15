package middlewares

import (
	"movie_planet/models"
	"movie_planet/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ParseToken(tokenStr string) (*models.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.UserClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(os.Getenv("SECRET_TOKEN")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.UserClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			util.ErrorJson(c, http.StatusUnauthorized, "Not login")
			c.Abort()
			return
		}
		user, err := ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid token."+err.Error())
			c.Abort()
			return
		}
		c.Set("email", user.Email)
		c.Next()
	}
}
