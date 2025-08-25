package routes

import (
	"github.com/SulthanAndyno/url-shortener/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/shorten", controllers.ShortenURL)
	router.GET("/:short_code", controllers.RedirectURL)
	router.GET("/stats/:short_code", controllers.GetURLStats)
	
	return router
}