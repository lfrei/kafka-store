package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func postProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func startWebServer(endpoint string) {
	r := mux.NewRouter()
	r.HandleFunc(endpoint, getProducts).Methods(http.MethodGet)
	r.HandleFunc(endpoint, postProducts).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	startWebServer("/products")
}
