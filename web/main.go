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

	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/test-api", HelloApi)
	err := http.ListenAndServe(":8000", nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

	log.Print("Completed.")
}

func HelloWorld(w http.ResponseWriter, _ *http.Request) {
	message := os.Getenv("MESSAGE_ENV_VAR")
	secret := os.Getenv("SECRET_ENV_VAR")

	_, _ = fmt.Fprintf(w, "Hello world!!!\n")
	_, _ = fmt.Fprintf(w, "Message: %s\n", message)
	_, _ = fmt.Fprintf(w, "Secret: %s\n", secret)

	log.Print("Handled request")
}

func HelloApi(w http.ResponseWriter, _ *http.Request) {
	apiUrl := os.Getenv("API_URL")

	response, err := http.Get(apiUrl + "/api")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}
	_, _ = fmt.Fprintf(w, "Fetched data from %s\n", apiUrl)

	body := make([]byte, response.ContentLength)
	_, _ = response.Body.Read(body)

	var responseBody Response
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = fmt.Fprintf(w, "Hello world!!!\n")
	_, _ = fmt.Fprintf(w, responseBody.Message)
	log.Print("Handled request")
}
