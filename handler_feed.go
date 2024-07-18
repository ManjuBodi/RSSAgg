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

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("Couldn't create feed %v:", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("Couldn't get feeds %v:", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}
