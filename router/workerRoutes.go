package router

import (
	"project/controller"

	"github.com/gin-gonic/gin"
)

func WorkerRoutes(rg *gin.RouterGroup) {
	worker := rg.Group("/")
	worker.GET("/", controller.GetWorkers)
}
