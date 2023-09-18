package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	redis "github.com/omfj/lol/internal"
	h "github.com/omfj/lol/internal/handlers"
)

func main() {
	redis := &redis.Redis{
		Addr: "localhost:6379",
	}

	err := redis.Connect()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", h.HealthCheckHandler).Methods("GET")
	r.HandleFunc("/status", h.GetStatusHandler).Methods("GET")
	r.HandleFunc("/status", h.UpdateStatusHandler).Methods("POST")
	http.Handle("/", r)

	log.Default().Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
