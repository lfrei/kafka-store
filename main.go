package main

import (
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"time"
)

var products *cache.Cache

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	product, found := products.Get(productId)
	if found {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(product.(string)))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func postProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func startWebServer() {
	r := mux.NewRouter()
	r.HandleFunc("/product/{id}", getProduct).Methods(http.MethodGet)
	r.HandleFunc("/products", getProducts).Methods(http.MethodGet)
	r.HandleFunc("/products", postProducts).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func setupCache() {
	products = cache.New(5*time.Minute, 10*time.Minute)

	//add some test data
	products.Set("001", "IPhone", cache.NoExpiration)
	products.Set("002", "IPad", cache.NoExpiration)
}

func main() {
	setupCache()
	startWebServer()
}
