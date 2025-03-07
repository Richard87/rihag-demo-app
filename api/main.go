package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	log.Print("Starting listening on http://localhost:8001/api...")

	http.HandleFunc("/api", HelloAPI)
	err := http.ListenAndServe(":8001", nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

	log.Print("Completed.")
}

func HelloAPI(response http.ResponseWriter, _ *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	content, _ := json.Marshal(Response{Message: "hello from API!"})
	response.Write(content)

	log.Print("Handled request")
}
