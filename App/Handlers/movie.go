package handlers

import (
	models "main/Models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// serializer
var movies = []models.Movie{
	{ID: "1", Title: "The Godfather", Director: "Francis Ford Coppola", Genre: "Drama", Description: "The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.", Year: 1972, Calification: 9.2, Duration: 175},
	{ID: "2", Title: "The Godfather: Part II", Director: "Francis Ford Coppola", Genre: "Drama", Description: "The early life and career of Vito Corleone in 1920s New York City is portrayed, while his son, Michael, expands and tightens his grip on the family crime syndicate.", Year: 1974, Calification: 9.0, Duration: 202},
	{ID: "3", Title: "The Dark Knight", Director: "Christopher Nolan", Genre: "Action", Description: "When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept one of the greatest psychological and physical tests of his ability to fight injustice.", Year: 2008, Calification: 9.0, Duration: 152},
}

func GetMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

func GetMovieByID(c *gin.Context) {

	id := c.Param("id")

	for _, a := range movies {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})

}