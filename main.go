package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := ConnectToRedis()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", HealthCheckHandler).Methods("GET")
	r.HandleFunc("/status", GetStatusHandler).Methods("GET")
	r.HandleFunc("/status", UpdateStatusHandler).Methods("POST")
	http.Handle("/", r)

	log.Default().Println("Listening on port http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
