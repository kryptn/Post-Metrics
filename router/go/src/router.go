package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"net/http"
	"encoding/json"
	
	"github.com/gorilla/mux"
)

type Metric struct{
	Metric	string
	Metrics []string
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/graphite", Graphite)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}

func Graphite(w http.ResponseWriter, r *http.Request) {
	var m Metric
	if r.Body == nil {
		fmt.Fprintln(w, "")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if m.Metric != "" {
		m.Metrics = append(m.Metrics, m.Metric)
	}

	fmt.Println(strings.Join(m.Metrics, "\n"))
	go SendUDP(strings.Join(m.Metrics, "\n"))

}

func SendUDP(s string) {
	conn, err := net.Dial("udp", "statsite:8125")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("trying to send:", s)
	conn.Write([]byte(s))
}

