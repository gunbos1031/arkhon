package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"net/http"
)

func apiHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Coin": "Arkhon",
	})
}

func Start() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../client/build", true)))
	
	api := router.Group("/api")
	{
		api.GET("/", apiHome)
	}
	
	router.Run(":80")
}