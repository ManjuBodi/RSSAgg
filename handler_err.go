package main

import "net/http"

//a very specific function specification you have to use if you want to define an http handler in a way that a go standard library expects

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithERROR(w, 400, "something went wrong")
}
