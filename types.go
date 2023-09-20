package main

type StatusResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type StatusPayload struct {
	Status int `json:"status"`
}
