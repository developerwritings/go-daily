package routes

import (
	"go_gin_crud/handlers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()
	router.POST("/", handlers.GetUser)
	return router
}
