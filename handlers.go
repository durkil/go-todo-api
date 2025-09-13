package main

import (
	"encoding/json"
	"net/http"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetTodos(w http.ResponseWriter, r *http.Request) {
	checkMethod(w, r, http.MethodGet)
	var todos []Todo
	DB.Find(&todos)
	sendJSONResponse(w, http.StatusCreated, todos)
}

func CreateTodos(w http.ResponseWriter, r *http.Request) {
	checkMethod(w, r, http.MethodPost)
	var newTodo Todo

	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	DB.Create(&newTodo)

	sendJSONResponse(w, http.StatusCreated, newTodo)
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	checkMethod(w, r, http.MethodGet)
	id := parseId(w, r)

	var todo Todo

	result := DB.Find(&todo, id)

	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sendJSONResponse(w, http.StatusCreated, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	checkMethod(w, r, http.MethodPatch)
	id := parseId(w, r)

	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var todo Todo
	result := DB.First(&todo, id)

	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	DB.Model(&todo).Updates(&updatedTodo)
	sendJSONResponse(w, http.StatusCreated, todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	checkMethod(w, r, http.MethodDelete)
	id := parseId(w, r)

	result := DB.Delete(&Todo{}, id)
	if result != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
