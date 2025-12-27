package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// ========== CRUD ========== //

type User struct {
	Id         uuid.UUID `json:"id"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
	Email      string    `json:"email"`
}

func (cfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}
	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	if err := decoder.Decode(&params); err != nil || params.Email == "" {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.db.CreateUser(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}
	respondWithJSON(w, http.StatusCreated, response{
		User: User{
			Id:         user.ID,
			Created_At: user.CreatedAt,
			Updated_At: user.UpdatedAt,
			Email:      user.Email,
		},
	})
}
