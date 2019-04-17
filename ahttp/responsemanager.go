package ahttp

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

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

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
