package application

import (
	"encoding/json"
	"net/http"

	model "github.com/wlanboy/gowebarduino/model"
)

/*WriteJSONErrorResponse with content type and status code*/
func WriteJSONErrorResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

/*WriteJSONResponse with content type and status code*/
func WriteJSONResponse(w http.ResponseWriter, command *model.Command, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(command)
}
