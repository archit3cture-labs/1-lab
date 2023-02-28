package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type JSONtime struct {
  CURRENT_TIME string `json:"time"`
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("got request\n")
  currentTime := JSONtime{CURRENT_TIME: time.Now().Format(time.RFC3339)}
  json_data, err := json.Marshal(currentTime)

	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(w, string(json_data))
}

func main() {}