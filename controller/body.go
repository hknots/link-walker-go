package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type RequestBody struct {
	URL string `json:"url"` 
}

func GetUrlFromBody(request *http.Request) (string, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return "", errors.New("failed to read request body")
	}
	defer request.Body.Close()

	var rb RequestBody
	err = json.Unmarshal(body, &rb)
	if err != nil {
		return "", errors.New("failed to decode request body")
	}

	if rb.URL == "" {
		return "", errors.New("url is missing in request body")
	}

	return rb.URL, nil
}
