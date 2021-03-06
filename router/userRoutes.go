package router

import (
	"project/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/")
	user.GET("/login", controller.ViewLogin)
	user.POST("/login", controller.GetUserByEmailAndPassword)
	user.GET("/new", controller.NewUser)
	user.POST("/", controller.AddUser)
	user.GET("/profile", controller.ViewProfile)
	user.PUT("/", controller.UpdateUser)
	user.DELETE("/:id", controller.DeleteUser)
	user.GET("/user/:id", controller.GetUserbyId)
}
