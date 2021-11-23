package main

import (
	"project/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	user := server.Group("/users")
	router.UserRoutes(user)
	worker := server.Group("/workers")
	router.WorkerRoutes(worker)

	server.Run(":4200")

}
