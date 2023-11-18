package main

import (
	handlers "main/Handlers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	//routers
	r.GET("/movies", handlers.GetMovies)
	r.GET("movies/:id", handlers.GetMovieByID)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8000")
}
