package main

import (
	"project/entities"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users", entities.GetUsers)
	r.POST("/users", entities.AddUser)
	r.GET("/users/:email/:password", entities.GetUserByEmailAndPassword)
	r.PUT("/users", entities.UpdateUser)
	r.Run()
}
