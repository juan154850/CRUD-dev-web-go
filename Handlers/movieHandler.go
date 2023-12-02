package Handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/juan154850/App/Controllers"
)

func GetMovies(c *gin.Context) {
	Controllers.GetMovies(c)
}

func GetMovieByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}
	Controllers.GetMovieByID(c, id)
}

func CreateMovie(c *gin.Context) {
	Controllers.CreateMovie(c)
}

func UpdateMovie(c *gin.Context) {
	Controllers.UpdateMovie(c)
}

func DeleteMovie(c *gin.Context) {
	Controllers.DeleteMovie(c)
}