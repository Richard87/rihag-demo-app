package main

import (
	"fmt"
	"io"
	"net/http"
)

func sendRequest(url string) (string, error) {
	apiResponse, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error calling API: %w", err)
	}
	defer apiResponse.Body.Close()

	body, err := io.ReadAll(apiResponse.Body)
	if err != nil {
		return "", fmt.Errorf("error reading API response: %w", err)
	}

	return string(body), nil
}
