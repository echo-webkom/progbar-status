package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Check redis connection
	err := RDB.Ping().Err()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("redis connection failed"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Get current status
	status, err := GetStatus()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to get status"))
		return
	}

	message := getMessage(status)

	statusStruct := StatusResponse{Status: status, Message: message}
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
	payload := StatusPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid payload"))
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
	err = ValidateToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid token"))
		return
	}

	// Set to desired status
	err = SetStatus(payload.Status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to set status"))
		return
	}

	// Get updated status
	status, err := GetStatus()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to get status"))
		return
	}

	message := getMessage(status)
	statusStruct := StatusResponse{Status: status, Message: message}
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

func getMessage(status int) string {
	switch status {
	case 0:
		return "Programmerbar er stengt."
	case 1:
		return "Programmerbar er åpen."
	default:
		return "Ukjent status"
	}
}
