package main

import (
	"github.com/doutnus/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", service.GetAlbums)
	router.GET("/albums/:id", service.GetSpecificyAlbum)

	router.POST("/albums", service.SetAlbums)
	router.PUT("/albums/:id", service.UpdateAlbum)
	router.DELETE("/albums/:id", service.DeleteAlbum)

	router.Run("127.0.0.1:8080")
}
