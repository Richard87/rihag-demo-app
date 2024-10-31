package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	log.Logger = log.Output(zerolog.NewConsoleWriter())
	log.Info().Msg("Starting listening on http://localhost:8000/...")
	log.Info().Msg("Starting listening on http://localhost:8000/test-api...")

	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/test-api", HelloApi)
	err := http.ListenAndServe(":8000", nil)

	log.Err(err).Msg("Completed.")
}

func HelloWorld(w http.ResponseWriter, _ *http.Request) {
	// message = os.Getenv("MESSAGE_ENV_VAR")
	// secret := os.Getenv("SECRET_ENV_VAR")

	_, _ = fmt.Fprintf(w, "Hello world!!!\n")
	// _, _ = fmt.Fprintf(w, "Message: %s\n", message)
	// _, _ = fmt.Fprintf(w, "Secret: %s\n", secret)

	log.Info().Msg("Handled request")
}

func HelloApi(w http.ResponseWriter, _ *http.Request) {
	log.Info().Msg("Handled request")
	apiUrl := os.Getenv("API_URL")

	response, err := http.Get(apiUrl + "/api")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

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
}
