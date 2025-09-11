package main

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json: "completed"`
}

var todos []Todo

var currentID int
