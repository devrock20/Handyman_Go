package router

import (
	"project/controller"

	"github.com/gin-gonic/gin"
)

func MainRoutes(rg *gin.RouterGroup) {
	site := rg.Group("/")
	//user.GET("/", controller.GetUsers)
	site.GET("/", controller.ViewIndex)

}
