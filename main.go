package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"web-service-gin/controllers"
)

func main() {

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	router.POST("/painting/create", controllers.AddPainting)

	router.GET("/artist/:artist", controllers.GetPaintingsByArtist)
	router.GET("/painting", controllers.GetPaintings)
	router.GET("/painting/:id/", controllers.GetPaintingByID)

	router.PUT("/artist/update/:id", controllers.UpdateArtist)
	router.PUT("/painting/update/:id", controllers.UpdatePainting)

	router.DELETE("/painting/delete/:id", controllers.DeleteOrder)

	router.Run(host + ":" + port)
}
