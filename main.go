package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Post struct {
	// `json:"user"` is for the json parsing of this User field. Otherwise, by default it's 'User'.
	User     string `json:"user"`
	Message  string  `json:"message"`
	Location Location `json:"location"`
}



func main() {
	fmt.Println("started-service")
	http.HandleFunc("/post", handlerPost)
	http.HandleFunc("/search", handlerSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a request")
	decoder := json.NewDecoder(r.Body)

	var p Post
	if err := decoder.Decode(&p); err != nil {
		panic(err)
		return
	}
	fmt.Fprintf(w, "Post is received : %s\n", p.Message)
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one request for search")
	lat := r.URL.Query().Get("lat");
	lon := r.URL.Query().Get("lon");


	fmt.Fprintf(w, "Lat: %s, Long: %s\n", lat, lon)
}