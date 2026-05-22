package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, status, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := JSONResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}
