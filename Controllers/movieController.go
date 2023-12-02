package Controllers

// import (
//     "encoding/json"
//     "net/http"
// 	Models "github.com/juan154850/App/Models"
// 	repository "github.com/juan154850/App/Repository"
// )


// var (
//     updateQuery = "UPDATE movies SET title=:title, director=:director, genre=:genre, description=:description, year=:year, calification=:calification, duration=:duration WHERE id=:id;"
//     deleteQuery = "DELETE FROM movies WHERE id=$1;"
//     selectQuery = "SELECT id, title, director, genre, description, year, calification, duration FROM movies WHERE id=$1;"
//     listQuery   = "SELECT id, title, director, genre, description, year, calification, duration FROM movies LIMIT $1 OFFSET $2;"
//     createQuery = "INSERT INTO movies (id, title, director, genre, description, year, calification, duration) VALUES (:id, :title, :director, :genre, :description, :year, :calification, :duration) RETURNING id;"
// )

// type Controller struct{
// 	repo repository.Repository
// }
// func NewController(repo repository.Repository) *Controller{
// 	return &Controller{repo: repo}
// }


// func (c *Controller) CreateMovie(w http.ResponseWriter, r *http.Request)(int,error){
// 	var newMovie Models.Movie
// 	err := json.NewDecoder(r.Body).Decode(&newMovie)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return 0, err 
// 	}

// 	id, err = c.repo.CreateMovie(newMovie)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return 0, err 
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	return id, nil
// }


// func (c *Controller) ReadMovie(w http.ResponseWriter, r *http.Request)([]byte,error){
// 	movie, err := c.repo.ReadMovie()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return nil, err
// 	}
// 	movieJson, err := json.Marshal(item)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return nil, err
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	return movieJson, nil
// }

// func (c *Controller) ListMovies(w http.ResponseWriter, r *http.Request, limit, offset int)([]byte,error){
// 	movies, err := c.repo.ListMovies(limit,offset)
// 	if err != nil{
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return nil, err 
// 	}
// 	moviesJson, err := json.Marshal(movies)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return nil, err 
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	return moviesJson, nil
// }

// func (c *Controller) UpdateMovie(w http.ResponseWriter, r *http.Request, id int) error {
// 	var updateMovie Models.Item
// 	err := json.NewDecoder(r.Body).Decode(&updateMovie)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}
// 	err = c.repo.UpdateMovie(updateMovie, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return err
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	return nil
// }

// func (c *Controller) DeleteMovie(w http.ResponseWriter, r *http.Request, id int) error {
// 	err := c.repo.DeleteMovie(id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return err
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	return nil
// }
