package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kargo-api/handler"
)

func main() {
	handler := handler.NewHandler()

	router := mux.NewRouter()

	router.HandleFunc("/kargo/getJobs", handler.GetSortedJobsHandler).Methods("GET")
	router.HandleFunc("/kargo/getBids", handler.GetSortedBidsHandler).Methods("GET")

	log.Println("Listening at port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
