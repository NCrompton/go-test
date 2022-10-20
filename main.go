package main

import (
	"fmt"
	/* "net/http" */
	"newpj/note/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("The program is running")
	router := gin.Default()
	router.GET("/user", controller.someNum())
	fmt.Print(random())

	router.Run("localhost:5000")
}
