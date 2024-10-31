package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	message := "Hello Radix!!!"
	secret := "hardcoded secret"

	// secret = os.Getenv("SECRET_ENV_VAR")
	// message = os.Getenv("MESSAGE_ENV_VAR")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello World"))
	_, _ = w.Write([]byte(message))
	_, _ = w.Write([]byte(secret))

	log.Info().Msg("Handled request")
}

func HelloApi(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Handled request")
	hostname := "localhost:8001"
	// hostname = os.Getenv("API_HOSTNAME")

	response, err := http.Get(fmt.Sprintf("http://%s/api", hostname))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	body := make([]byte, 1024)
	length, _ := response.Body.Read(body)

	var responseBody Response
	err = json.Unmarshal(body[:length], &responseBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write([]byte(responseBody.Message))
}
