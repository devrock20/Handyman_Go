package main

import (
	"project/router"

	method "github.com/bu/gin-method-override"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	gv := goview.Config{
		Root:         "views",
		Extension:    ".tmpl",
		Partials:     []string{"partials/header", "partials/footer"},
		DisableCache: true,
	}
	server.Use(method.ProcessMethodOverride(server))

	//Set new instance

	server.HTMLRender = ginview.New(gv)
	server.Static("/css", "public/css")
	site := server.Group("/")
	router.MainRoutes(site)
	user := server.Group("/users")
	router.UserRoutes(user)
	worker := server.Group("/workers")
	router.WorkerRoutes(worker)

	server.Run(":4200")

}
