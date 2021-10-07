package handlers

import (
	"Downloads/MovieDbProj/entities"
	"Downloads/MovieDbProj/service"
	"encoding/json"
	"net/http"
	"log"
	//"fmt"
	"github.com/gorilla/mux"

)

//Our MovieHandler struct will perform the service given to it by our service layer 
type MovieHandler struct {
	PerformService service.Service

}


func NewMovieHandler(s service.Service) MovieHandler{
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

	err = mvfile.PerformService.Repo.NewMovie(postResult)
	if err != nil {
		log.Fatalln(err)
	}
	


	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("Movie added successfully!"))

}



//A method on on our movie handler struct that will talk to our server and handle the request and our response
//Our path contains a variable (the id) so we pass in the request (which will be our movie/"id") &
//Set the key of that to "id" 
//Line 68 accesses our movie handler --> accesses our service layer ---> accesses our FindById in our repo layer
//And again returns the movie
//Now that we have our movie we convert it back to a json object (line 75)
//Lines 82-84 access our ResponseWriter and serves getMovie

func (mvfile MovieHandler)GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	getResult, err := mvfile.PerformService.Repo.FindById(id)
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

	err := mvfile.PerformService.Repo.DeleteById(id)
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

	//Unmarshal r.Body into movie and then pass what i've unmarshaled to UpdateById
    
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatalln(err)
	}
	

	err = mvfile.PerformService.Repo.UpdateById(id, movie)
	if err != nil {
		log.Fatalln(err)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Movie successfully updated in database!"))
}

