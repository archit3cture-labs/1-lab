package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
    )

type JSONtime struct {
  CURRENT_TIME string json:"time"
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("got request\n")
  currentTime := JSONtime{CURRENT_TIME: time.Now().Format(time.RFC3339)}
  json_data, err := json.Marshal(currentTime)
}

func main() {}