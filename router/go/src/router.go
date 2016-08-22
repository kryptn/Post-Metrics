package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/graphite", Graphite)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome hi")
}

func Graphite(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Grahpite page")
}

