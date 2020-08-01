package main

import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status int `json:"status"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		payload, _ := json.Marshal(Response{"Welcome to this is awesome API.", http.StatusOK})

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(payload))

	})

	fmt.Println("Server listening!")
	log.Fatal(http.ListenAndServe(":8000", r))
}
