package handlers 

import (
	"github.com/gorilla/mux"
	//"net/http"
)



func NewRouter(handler MovieHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movies/{Id}", handler.PutUpdateMovie).Methods("PUT")
	r.HandleFunc("/movies", handler.PostNewMovie).Methods("POST")
	r.HandleFunc("/movies/{Id}", handler.GetMovieById).Methods("GET")
	r.HandleFunc("/movies/{Id}", handler.DeleteMovie).Methods("DELETE")
	

	return r

}



