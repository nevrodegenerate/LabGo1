package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().UTC().Format(time.RFC3339)
	response := TimeResponse{Time: currentTime}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/time", handler)
	log.Fatal(http.ListenAndServe(":8795", nil))
}
