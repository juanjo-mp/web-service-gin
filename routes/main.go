package routes

import (
	"github.com/gin-gonic/gin"

	"web-service-gin/controllers"
)

func Painting(g *gin.RouterGroup) {
	g.POST("/create", controllers.AddPainting)
	g.GET("/", controllers.GetPaintings)
	g.GET("/:id", controllers.GetPaintingByID)
	g.PUT("/update/:id", controllers.UpdatePainting)
	g.DELETE("/delete/:id", controllers.DeleteOrder)
}

func Artist(g *gin.RouterGroup) {
	g.GET("/:artist", controllers.GetPaintingsByArtist)
	g.PUT("/update/:id", controllers.UpdateArtist)
}
