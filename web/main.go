package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	log.Print("Starting listening on http://localhost:8000/...")
	log.Print("Starting listening on http://localhost:8000/test-api...")

	http.HandleFunc("/", HelloWorldHandler)
	http.HandleFunc("/test-api", HelloApiHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

	log.Print("Completed.")
}

func HelloWorldHandler(response http.ResponseWriter, _ *http.Request) {
	message := os.Getenv("MESSAGE_ENV_VAR")
	secret := os.Getenv("SECRET_ENV_VAR")

	response.Write([]byte(fmt.Sprintf("Hello world!!!\n")))
	response.Write([]byte(fmt.Sprintf("Message: %s\n", message)))
	response.Write([]byte(fmt.Sprintf("Secret: %s\n", secret)))

	log.Print("Handled request")
}

func HelloApiHandler(response http.ResponseWriter, _ *http.Request) {
	apiUrl := os.Getenv("API_URL")

	apiResponse, err := http.Get(apiUrl + "/api")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	response.Write([]byte(fmt.Sprintf("Fetched data from %s\n", apiUrl)))

	body := make([]byte, apiResponse.ContentLength)
	_, _ = apiResponse.Body.Read(body)

	var responseBody Response
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	response.Write([]byte(fmt.Sprintf("Hello world!!!\n")))
	response.Write([]byte(fmt.Sprintf("Response from API: %s\n", responseBody.Message)))
	log.Print("Handled request")
}
