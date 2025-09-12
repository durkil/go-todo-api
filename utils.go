package main

import (
	"net/http"
	"strconv"
	"strings"
	"encoding/json"
)

func checkMethod(w http.ResponseWriter, r *http.Request, expectedMethod string) {
	if r.Method != expectedMethod {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func parseId(w http.ResponseWriter, r *http.Request) int {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return 0
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	return id
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, encodedItem interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(encodedItem)
}