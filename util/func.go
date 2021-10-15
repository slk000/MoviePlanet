package util

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestRequest(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Print(string(jsonData))
	c.JSON(http.StatusOK, "ok")
}
