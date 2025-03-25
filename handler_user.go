package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZaxVaxZ/RSSFeedBackend/internal/auth"
	"github.com/ZaxVaxZ/RSSFeedBackend/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"username"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error parsing JSON", err))
		return
	} else if len(params.Name) < 3 {
		respondWithError(w, http.StatusBadRequest, "Error: Username must be at least 3 characters")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:       uuid.New(),
		Username: params.Name,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Couldn't create user: ", err))
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header) 

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error:", err))
		return
	}
	
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Couldn't get user: ", err))
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (apiCfg *apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header) 

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error:", err))
		return
	}
	
	_, err = apiCfg.DB.DeleteUserByAPIKey(r.Context(), apiKey)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Couldn't delete user: ", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
