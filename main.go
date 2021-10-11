package main

import (
	"Downloads/MovieDbProj/handlers"
	"log"
	//"fmt"
	"Downloads/MovieDbProj/repo"
	//"Downloads/MovieDbProj/entities"
	"Downloads/MovieDbProj/service"
	"net/http"
	"path/filepath"
)

func main() {

	ourJsonFile := "moviedb.json"

	extension := filepath.Ext(ourJsonFile)

	if extension != ".json" {
		log.Fatalln("File extension incorrect.")
	}

	repoLayer := repo.NewRepo(ourJsonFile)

	serviceLayer := service.NewService(repoLayer)

	handlersLayer:= handlers.NewMovieHandler(serviceLayer)

	router := handlers.NewRouter(handlersLayer)

	svr := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	log.Fatalln(svr.ListenAndServe())

}

//server := handlers.NewServer()

//log.Fatal(server.ListenAndServe())
