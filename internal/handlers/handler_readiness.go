package handlers

import (
	"net/http"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	RespondWithJson(w, 200, struct{}{})
}

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, 400, "Something went wrong")
}
