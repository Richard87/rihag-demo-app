package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	log.Logger = log.Output(zerolog.NewConsoleWriter())
	log.Info().Msg("Starting listening on http://localhost:8001/api...")

	http.HandleFunc("/api", HelloAPI)
	err := http.ListenAndServe(":8001", nil)

	log.Err(err).Msg("Completed.")
}

func HelloAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	content, _ := json.Marshal(Response{Message: "hello from API!"})
	_, _ = w.Write(content)

	log.Info().Msg("Handled request")
}
