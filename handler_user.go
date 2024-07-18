package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ManjuBodi/RSSAgg/internal/database"
	"github.com/google/uuid"
)

//a very specific function specification you have to use if you want to define an http handler in a way that a go standard library expects

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("Couldn't create user %v:", err))
		return
	}
	respondWithJSON(w, 201, databaseUsertoUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUsertoUser(user))
}

func (apiCfg *apiConfig) handlerPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("Couldn't get Posts %v:", err))
		return
	}
	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
