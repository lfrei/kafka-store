package main

import (
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var products *cache.Cache

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
	productId := mux.Vars(r)["id"]
	product, err := ioutil.ReadAll(r.Body)
	if err == nil {
		products.Set(productId, string(product), cache.NoExpiration)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "product created"}`))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func startWebServer() {
	r := mux.NewRouter()
	r.HandleFunc("/product/{id}", getProduct).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", postProducts).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func setupCache() {
	products = cache.New(5*time.Minute, 10*time.Minute)
}

func main() {
	setupCache()
	startWebServer()
}
