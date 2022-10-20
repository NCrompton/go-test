package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, user{Name: "CX", Age: 16})
}

func someNum() int {
	return 20
}
