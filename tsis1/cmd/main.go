package main

import (
	"log"
	"net/http"

	"github.com/balgabekj/go/tsis1/api"
	"github.com/gorilla/mux"
)

const port = ":8080"

func main() {
	log.Println("starting API server")
	router := mux.NewRouter()
	log.Println("creating routes")
	api.SetupRoutes(router)
	http.Handle("/",router)
	log.Println("Server started at port", port)

	http.ListenAndServe(port,router)
}