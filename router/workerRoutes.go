package router

import (
	"project/controller"

	"github.com/gin-gonic/gin"
)

func WorkerRoutes(rg *gin.RouterGroup) {
	worker := rg.Group("/")

	worker.POST("/", controller.AddWorker)
	worker.GET("/new", controller.ViewWorkerNew)
	worker.GET("/show", controller.GetAllWorkers)
	worker.GET("/worker/:id", controller.GetWorkerById)
	worker.PUT("/updateWorker", controller.UpdateWorker)
	worker.DELETE("/:id", controller.DeleteWorker)
	worker.GET("/login", controller.ViewWorkerLogin)
	worker.POST("/login", controller.WorkerLogin)
}
