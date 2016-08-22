package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	
	"github.com/gorilla/mux"
)

type Metric struct{
	Metric	string
}

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
	var m Metric
	if r.Body == nil {
		fmt.Fprintln(w, "Hay")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(m.Metric)

}

