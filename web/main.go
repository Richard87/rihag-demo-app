package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

var message = os.Getenv("MESSAGE_ENV_VAR")
var secret = os.Getenv("SECRET_ENV_VAR")
var apiUrl = os.Getenv("API_URL")

func main() {
	log.Printf("Starting...")
	log.Printf("Config: MESSAGE_ENV_VAR: %s", message)
	log.Printf("Config: SECRET_ENV_VAR: %s", secret)
	log.Printf("Config: API_URL: %s", apiUrl)
	log.Printf("Starting listening on http://localhost:8000/...")
	log.Printf("Starting listening on http://localhost:8000/test-api...")

	http.HandleFunc("/", HelloWorldHandler)
	http.HandleFunc("/test-api", HelloApiHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

	log.Print("Completed.")
}

func HelloWorldHandler(response http.ResponseWriter, _ *http.Request) {

	response.Write([]byte(fmt.Sprintf("Hello world!!!\n")))
	response.Write([]byte(fmt.Sprintf("Message: %s\n", message)))
	response.Write([]byte(fmt.Sprintf("Secret: %s\n", secret)))
	response.Write([]byte(fmt.Sprintf("\n\nTest API: /api%s\n", secret)))

	log.Print("Handled request")
}

func HelloApiHandler(response http.ResponseWriter, _ *http.Request) {
	responseBody, err := sendRequest(apiUrl + "/api")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		log.Printf("Error: %v", err.Error())
		return
	}

	response.Write([]byte(fmt.Sprintf("Fetched data from %s\n", apiUrl)))
	response.Write([]byte(fmt.Sprintf("Hello world!!!\n")))
	response.Write([]byte(fmt.Sprintf("Response from API: %s\n", responseBody.Message)))
	log.Print("Handled request")
}
