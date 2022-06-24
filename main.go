package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/notes", CreateNote).Methods(http.MethodPost)  // adding a note
	r.HandleFunc("/notes", GetNoteByTag).Queries("tag", "{tag}") // getting notes with the "<specific>" tag
	r.HandleFunc("/notes", GetAllNotes)                          // getting all notes
	log.Fatal(http.ListenAndServe(":4000", r))
}
