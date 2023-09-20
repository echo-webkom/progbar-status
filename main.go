package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	redis := &Redis{
		Addr: "localhost:6379",
	}

	err := redis.Connect()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", HealthCheckHandler).Methods("GET")
	r.HandleFunc("/status", GetStatusHandler).Methods("GET")
	r.HandleFunc("/status", UpdateStatusHandler).Methods("POST")
	http.Handle("/", r)

	log.Default().Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
