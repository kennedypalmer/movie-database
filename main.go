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

	r := repo.NewRepo(ourJsonFile)

	svc := service.NewService(r)

	hdlr := handlers.NewMovieHandler(svc)

	router := handlers.NewRouter(hdlr)

	svr := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	log.Fatalln(svr.ListenAndServe())

}

//server := handlers.NewServer()

//log.Fatal(server.ListenAndServe())
