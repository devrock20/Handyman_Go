package main

import (
	"project/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")

	server.Static("/css", "public/css")
	site := server.Group("/")
	router.MainRoutes(site)
	user := server.Group("/users")
	router.UserRoutes(user)
	worker := server.Group("/workers")
	router.WorkerRoutes(worker)

	server.Run(":4200")

}
