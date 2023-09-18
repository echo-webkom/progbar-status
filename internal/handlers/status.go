package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/omfj/lol/internal/models"
	"github.com/omfj/lol/internal/security"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Get current status
	status, err := models.GetStatus()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to get status"))
		return
	}

	statusStruct := StatusResponse{Status: status}
	statusJson, err := json.Marshal(statusStruct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to marshal status"))
		return
	}

	// Return current status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(statusJson))
}

func UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	// Check if status is set
	if status == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("status query parameter is required"))
		return
	}

	// Get token from header
	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("authorization header is required"))
		return
	}

	// Remove "Bearer" from token and validate it
	token = strings.Split(token, " ")[1]
	err := security.ValidateToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid token"))
		return
	}

	// Set to desired status
	err = models.SetStatus(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to set status"))
		return
	}

	// Get updated status
	status, err = models.GetStatus()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to get status"))
		return
	}

	statusStruct := StatusResponse{Status: status}
	statusJson, err := json.Marshal(statusStruct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to marshal status"))
		return
	}

	// Return updated status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(statusJson))
}
