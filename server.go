package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	CurrentTime string `json:"currentTime"`
}

func timeHandler(responseWriter http.ResponseWriter, request *http.Request) {
	currentTimeUTC := time.Now().UTC().Format(time.RFC3339)
	timeResponse := TimeResponse{CurrentTime: currentTimeUTC}

	responseWriter.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(responseWriter).Encode(timeResponse)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/time", timeHandler)
	log.Fatal(http.ListenAndServe(":8795", nil))
}
