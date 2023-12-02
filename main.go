package main

import (
	_ "fmt"

	"github.com/gin-gonic/gin"
	handlers "github.com/juan154850/App/Handlers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	//routers
	r.GET("/dada", handlers.GetMovies)
	r.GET("/movies", handlers.GetMovies)
	r.GET("movies/:id", handlers.GetMovieByID)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8000")
}
