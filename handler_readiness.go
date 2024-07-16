package main

import "net/http"

//a very specific function specification you have to use if you want to define an http handler in a way that a go standard library expects

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
