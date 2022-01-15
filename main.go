package main

import (
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"web-service-gin/routes"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	v1 := r.Group("/api")

	routes.Painting(v1.Group("/painting"))
	routes.Artist(v1.Group("/artist"))

	r.Run(host + ":" + port)
}
