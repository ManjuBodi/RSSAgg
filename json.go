package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithERROR(w http.ResponseWriter, code int, msg string) { //this function, instead of taking a json will take a message string and this function is going to going to format that message into a consistent JSON object every single time
	if code > 499 { //errors in 400 range are client side errors, so we dont really know about them, they might not be using it well
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json: "error"`
	}
	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal Json response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
