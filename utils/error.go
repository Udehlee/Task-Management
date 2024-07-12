package utils

import (
	"encoding/json"
	"net/http"
)

func UnsucessfulRequest(w http.ResponseWriter, status, message string, statuscode int) {

	regErr := map[string]interface{}{
		"status":     status,
		"message":    message,
		"statusCode": statuscode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(regErr)
}
