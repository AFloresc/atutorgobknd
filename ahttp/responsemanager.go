package ahttp

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

//COMMON METHODS:

//RespondWithError : error responder in JSON format
// func RespondWithError(w http.ResponseWriter, code int, message string) {
// 	RespondWithJSON(w, code, map[string]string{"error": message})
// }

//RespondWithError : error responder in JSON format
func RespondWithError(w http.ResponseWriter, code int, message Error) {
	RespondWithJSON(w, code, map[string]string{"error": message.Message})
}

//RespondWithJSON : JSON format
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//END COMMON METHODS
