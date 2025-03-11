package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendRequest(url string) (*Response, error) {
	apiResponse, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error calling API: %w", err)
	}
	defer apiResponse.Body.Close()

	body := make([]byte, apiResponse.ContentLength)
	_, err = apiResponse.Body.Read(body)
	if err != nil {
		return nil, fmt.Errorf("error reading API response: %w", err)
	}

	var responseBody Response
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return nil, fmt.Errorf("error parsing API response: %w", err)
	}

	return &responseBody, nil
}
