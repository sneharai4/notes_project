package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Adding few sample notes for get calls. This need not be present when notes are saved in in-memory cache or db.
var m1 = "Clean dishes"
var message1 = &m1

var m2 = "Buy chips and salsa"
var message2 = &m2

var m3 = "Buy apples and oranges"
var message3 = &m3

var m4 = "Feed the cat"
var message4 = &m4

var t1 = "chores"
var tag1 = &t1

var t2 = "groceries"
var tag2 = &t2

var t3 = "groceries"
var tag3 = &t3

var notes = Notes{
	{Message: message1, Tag: tag1},
	{Message: message2, Tag: tag2},
	{Message: message3, Tag: tag3},
	{Message: message4},
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all notes called")
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

func GetNoteByTag(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")
	fmt.Printf("get note by tag %s called\n", tag)

	var noteByTag Notes
	for _, note := range notes {
		if note.Tag == &tag {
			noteByTag = append(noteByTag, note)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(noteByTag)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create note called")
	note := Note{}

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding response object", http.StatusBadRequest)
		return
	}

	if note.Message == nil || *note.Message == "" {
		http.Error(w, "message field is missing or is empty in request payload", http.StatusBadRequest)
		return
	}
	notes = append(notes, note)

	response, err := json.Marshal(&note)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
