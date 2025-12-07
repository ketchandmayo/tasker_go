package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

type Response struct {
	Word string `json:"word"`
}

var words = []string{
	"gopher", "programming", "golang", "developer",
	"software", "code", "computer", "algorithm",
	"network", "server",
}

func randomWord() string {
	return words[rand.Intn(len(words))]
}

func wordHandler(w http.ResponseWriter, r *http.Request) {
	word := randomWord()
	response := Response{
		Word: word,
	}

	log.Println(word)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/random-word", wordHandler)

	log.Println("Server started on http://localhost:8080/random-word")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
