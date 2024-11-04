package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var request requestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	task = request.Message
	fmt.Fprintf(w, "Task received", task)
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, task)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/post", PostHandler).Methods("POST")
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
