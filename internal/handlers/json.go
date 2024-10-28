package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", message)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJson(w, code, errResponse{message})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}
