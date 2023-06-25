package main

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		todos := Todo{
			Id:        1,
			Task:      "First todo item",
			Completed: false,
		}

		json.NewEncoder(w).Encode(todos)
	})

	http.ListenAndServe(":4000", nil)
}
