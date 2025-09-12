package main

import (
	"encoding/json"
	"net/http"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	checkMethod(w, r, http.MethodGet)
	sendJSONResponse(w, http.StatusCreated, todos)
}

func CreateTodos(w http.ResponseWriter, r *http.Request) {
	checkMethod(w, r, http.MethodPost)
	
	var newTodo Todo

	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	currentID++
	newTodo.ID = currentID

	todos = append(todos, newTodo)

	sendJSONResponse(w, http.StatusCreated, newTodo)
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	checkMethod(w, r, http.MethodGet)

	id := parseId(w, r)

	for _, todo := range todos {
		if todo.ID == id {
			sendJSONResponse(w, http.StatusCreated, todo)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	checkMethod(w, r, http.MethodPut)

	id := parseId(w, r)

	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i := range todos {
		if todos[i].ID == id {
			updatedTodo.ID = id
			todos[i] = updatedTodo
			sendJSONResponse(w, http.StatusCreated, updatedTodo)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	checkMethod(w, r, http.MethodDelete)

	id := parseId(w, r)

	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
