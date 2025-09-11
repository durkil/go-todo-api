package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimRight(r.URL.Path, "/")
		if path == "/todos" {
			switch r.Method {
			case http.MethodGet:
				GetTodos(w, r)
			case http.MethodPost:
				CreateTodos(w, r)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else {
			switch r.Method {
			case http.MethodGet:
				GetTodoByID(w, r)
			case http.MethodPut:
				UpdateTodo(w, r)
			case http.MethodDelete:
				DeleteTodo(w, r)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		}
	})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	todos = append(todos, Todo{
		ID:        1,
		Title:     "Learn Go",
		Completed: false,
	})

	todos = append(todos, Todo{
		ID:        2,
		Title:     "Learn Go",
		Completed: false,
	})

	todos = append(todos, Todo{
		ID:        3,
		Title:     "Learn Go",
		Completed: false,
	})

	todos = append(todos, Todo{
		ID:        4,
		Title:     "Learn Go",
		Completed: false,
	})
	currentID = 4
}
