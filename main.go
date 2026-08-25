package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type helloResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(helloResponse{Message: "Hello, World!"}); err != nil {
		log.Printf("encode response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	log.Println("listening on :5050")
	if err := http.ListenAndServe(":5050", mux); err != nil {
		log.Fatal(err)
	}
}
