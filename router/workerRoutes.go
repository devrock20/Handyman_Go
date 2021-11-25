package router

import (
	"project/controller"

	"github.com/gin-gonic/gin"
)

func WorkerRoutes(rg *gin.RouterGroup) {
	worker := rg.Group("/")

	worker.POST("/", controller.AddWorker)
	worker.GET("/", controller.GetAllWorkers)
	worker.GET("/worker/:id", controller.GetWorkerById)
	worker.PUT("/updateWorker", controller.UpdateWorker)
	worker.DELETE("/:id", controller.DeleteWorker)
	// worker.GET("/login", controller.GetUserByEmailAndPassword)

	worker.POST("/login", controller.WorkerLogin)
}
