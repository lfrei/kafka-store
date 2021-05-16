package main

import (
	"fmt"
	"log"
	"net/http"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
	// replace
	w.Write([]byte("Iphone,Ipad"))
}

func startWebServer(endpoint string, port int) {
	http.HandleFunc(endpoint, productHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}

func main() {
	startWebServer("/products", 8080)
}
