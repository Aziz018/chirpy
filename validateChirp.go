package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type ChirpRequest struct {
	Body string `json:"body"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type ValidResponse struct {
	Valid bool `json:"valid"`
}

type CleanedResponse struct {
	CleanedBody string `json:"cleaned_body"`
}

var profaneWords = []string{"kerfuffle", "sharbert", "fornax"}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := ChirpRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid JSON"})
		return
	}

	if len(req.Body) > 140 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Chirp is too long"})
		return
	}

	words := strings.Split(req.Body, " ")
	for i, word := range words {
		lowerWord := strings.ToLower(word)
		for _, profane := range profaneWords {
			if lowerWord == profane {
				words[i] = "****"
			}
		}
	}
	cleaned := strings.Join(words, " ")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CleanedResponse{CleanedBody: cleaned})
}
