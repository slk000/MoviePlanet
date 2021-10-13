package util

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorJson(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"error": data})
}

func SuccessJson(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"msg": data})
}
