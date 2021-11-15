package main

import (
	"project/entities"

	"github.com/gin-gonic/gin"

	"fmt"
)

func main() {
	fmt.Print("Modified Contents")
	r := gin.Default()
	r.GET("/users", entities.GetUsers)
	r.POST("/users", entities.AddUser)
	r.Run()
}
