package handlers

import (
	"Downloads/MovieDbProj/entities"
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"

)

type ServiceLayer interface {
    PostMovie(movie entities.Movie) error 
    FindMovieById(id string) (entities.Movie, error)
    DeleteMovieById(id string) error
    UpdateMovieById(id string, movie entities.Movie) error
}

type MovieHandler struct {
	PerformService ServiceLayer

}


func NewMovieHandler(s ServiceLayer) MovieHandler{
	return MovieHandler{
		PerformService : s, 
	}
}


func (mvfile MovieHandler) PostNewMovie(w http.ResponseWriter, r *http.Request) {
		postResult := entities.Movie{}


	err := json.NewDecoder(r.Body).Decode(&postResult)
	if err != nil {
		log.Fatalln(err)
	}

	err = mvfile.PerformService.PostMovie(postResult)
	if err != nil {
		log.Fatalln(err)
	}
	


	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("Movie added successfully!"))

}


func (mvfile MovieHandler)GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	getResult, err := mvfile.PerformService.FindMovieById(id)
	if err != nil {
		log.Fatalln(err)
	}

	getMovie, err := json.MarshalIndent(getResult, "", " ")
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(getMovie)




}

func (mvfile MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["Id"]

	err := mvfile.PerformService.DeleteMovieById(id)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Movie successfully deleted from database!"))


}


func (mvfile MovieHandler) PutUpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["Id"]

	movie := entities.Movie{}

    
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatalln(err)
	}
	

	err = mvfile.PerformService.UpdateMovieById(id, movie)
	if err != nil {
		log.Fatalln(err)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Movie successfully updated in database!"))
}

