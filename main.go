package main

import (
	"github.com/gin-gonic/gin"	
	Database "github.com/juan154850/App/Database"
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

	Database.InitDB()
	defer Database.CloseDB()

	r := setupRouter()
	r.Run(":8000")
}
