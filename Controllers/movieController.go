// Controllers/movieController.go
package Controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/juan154850/App/Database"
	"github.com/juan154850/App/Models"
	"net/http"
	"strconv"
)

func GetMovies(c *gin.Context) {
	rows, err := Database.Conn.Query(context.Background(), "SELECT * FROM movies")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var movies []Models.Movie
	for rows.Next() {
		var movie Models.Movie
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Director, &movie.Genre, &movie.Description, &movie.Year, &movie.Calification, &movie.Duration)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		movies = append(movies, movie)
	}

	c.JSON(http.StatusOK, movies)
}

func GetMovieByID(c *gin.Context, id int) {
	row := Database.Conn.QueryRow(context.Background(), "SELECT * FROM movies WHERE id = $1", id)
	var movie Models.Movie
	err := row.Scan(&movie.ID, &movie.Title, &movie.Director, &movie.Genre, &movie.Description, &movie.Year, &movie.Calification, &movie.Duration)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "No movie found with given id"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, movie)
}

func CreateMovie(c *gin.Context) {
	var movie Models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie.ID = 0 // Set ID to 0 to ensure the database generates a new ID

	row := Database.Conn.QueryRow(context.Background(), "SELECT 1 FROM movies WHERE title = $1 OR description = $2", movie.Title, movie.Description)
	var exists int
	err := row.Scan(&exists)
	if err != pgx.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Movie with this title or description already exists"})
		return
	}

	_, err = Database.Conn.Exec(context.Background(), "INSERT INTO movies (title, director, genre, description, year, calification, duration) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		movie.Title, movie.Director, movie.Genre, movie.Description, movie.Year, movie.Calification, movie.Duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Movie created successfully"})
}

func UpdateMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var movie Models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmdTag, err := Database.Conn.Exec(context.Background(), "UPDATE movies SET title = $1, director = $2, genre = $3, description = $4, year = $5, calification = $6, duration = $7 WHERE id = $8",
		movie.Title, movie.Director, movie.Genre, movie.Description, movie.Year, movie.Calification, movie.Duration, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if cmdTag.RowsAffected() > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "Movie updated successfully"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": "No movie found with given id"})
	}
}

func DeleteMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	cmdTag, err := Database.Conn.Exec(context.Background(), "DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if cmdTag.RowsAffected() > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "Movie deleted successfully"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": "No movie found with given id"})
	}
}