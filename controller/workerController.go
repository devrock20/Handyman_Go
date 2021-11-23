package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetWorkers(c *gin.Context) {
	fmt.Println(c)
}
