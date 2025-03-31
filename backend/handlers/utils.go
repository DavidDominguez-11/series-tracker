package handlers

import (
	"encoding/json"
	"net/http"

	"series-tracker-backend/models"
)

func respondWithJSON(w http.ResponseWriter, payload models.ApiResponse, status ...int) {
	w.Header().Set("Content-Type", "application/json")
	
	statusCode := http.StatusOK
	if len(status) > 0 {
		statusCode = status[0]
	}
	w.WriteHeader(statusCode)
	
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, message string, status int) {
	respondWithJSON(w, models.ApiResponse{
		Success: false,
		Message: message,
	}, status)
}