package main

type Note struct {
	Message *string `json:"message,required"`
	Tag     *string `json:"tag,omitempty"`
}

type Notes []Note
