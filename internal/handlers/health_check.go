package handlers

import (
	"net/http"

	redis "github.com/omfj/lol/internal"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Check redis connection
	err := redis.RDB.Ping().Err()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("redis connection failed"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
