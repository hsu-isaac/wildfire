package routes

import (
	"wildfire/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", handler.FetchNameAndJokeHandler)
}
