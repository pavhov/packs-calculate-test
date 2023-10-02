package api

import (
	"github.com/gin-gonic/gin"
)

func StartServer() error {
	r := gin.Default()
	r.POST("/calculate", CalculateHandler)
	return r.Run(":8080")
}
