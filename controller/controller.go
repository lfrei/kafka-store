package controller

import (
	"github.com/gorilla/mux"
	"github.com/lfrei/kafka-store/store"
	"io/ioutil"
	"log"
	"net/http"
)

func getProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	product := store.GetProduct(id)
	jsonResponse(w, http.StatusOK, product)
}

func postProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	product, err := ioutil.ReadAll(r.Body)
	if err == nil {
		store.AddProduct(id, string(product))
		jsonResponse(w, http.StatusCreated, `{"message": "product created"}`)
	} else {
		jsonResponse(w, http.StatusBadRequest, `{"message": "invalid product""}`)
	}
}

func jsonResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/product/{id}", getProduct).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", postProduct).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}
