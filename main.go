package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Write([]byte("You're visiting: " + path))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("View snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Create snippet"))
}

func main() {
	mux := http.NewServeMux()
  
  // Handler
	mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet/view", snippetView)
  mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("SERVER STARTED AT PORT 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
