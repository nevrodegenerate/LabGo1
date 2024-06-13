package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// TimeResponse представляет структуру ответа с текущим временем
type TimeResponse struct {
	CurrentTime string `json:"currentTime"`
}

// timeHandler обрабатывает HTTP-запросы и возвращает текущее время в формате JSON
func timeHandler(responseWriter http.ResponseWriter, request *http.Request) {
	currentTimeUTC := time.Now().UTC().Format(time.RFC3339)
	timeResponse := TimeResponse{CurrentTime: currentTimeUTC}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(timeResponse)
}

func main() {
	http.HandleFunc("/time", timeHandler)
	log.Fatal(http.ListenAndServe(":8795", nil))
}
