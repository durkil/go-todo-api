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


