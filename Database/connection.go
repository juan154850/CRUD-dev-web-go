// Database/database.go
package Database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func InitDB() {
	var err error
	Conn, err = pgx.Connect(context.Background(), "postgresql://devWebUser:devWebPassword@db:5432/crudDevWeb")
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}
	fmt.Println("Connected to database!")

	err = createTables()
	if err != nil {
		fmt.Println("Error creating tables: ", err)
		return
	}
	fmt.Println("Tables created successfully!")

	err = populateMovies()
	if err != nil {
		fmt.Println("Error populating movies: ", err)
		return
	}
	fmt.Println("Movies populated successfully!")
}

func createTables() error {
	_, err := Conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS movies (
			id SERIAL PRIMARY KEY,
			title VARCHAR(100),
			director VARCHAR(100),
			genre VARCHAR(100),
			description TEXT,
			year INT,
			calification REAL,
			duration INT
		)
	`)
	return err
}

func populateMovies() error {
	_, err := Conn.Exec(context.Background(), `
		INSERT INTO movies (id, title, director, genre, description, year, calification, duration) VALUES
		('1', 'The Godfather', 'Francis Ford Coppola', 'Drama', 'The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.', 1972, 9.2, 175),
		('2', 'The Godfather: Part II', 'Francis Ford Coppola', 'Drama', 'The early life and career of Vito Corleone in 1920s New York City is portrayed, while his son, Michael, expands and tightens his grip on the family crime syndicate.', 1974, 9.0, 202),
		('3', 'The Dark Knight', 'Christopher Nolan', 'Action', 'When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept one of the greatest psychological and physical tests of his ability to fight injustice.', 2008, 9.0, 152)
		ON CONFLICT (id) DO NOTHING
	`)
	return err
}

func CloseDB() {
	Conn.Close(context.Background())
}
