package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
}

func Fetcher() (Response, error) {
	var data Response
	var client = &http.Client{}

	request, err := http.NewRequest("GET", "https://mocki.io/v1/4459ceb7-791b-4948-87e1-db598b805587", nil)
	if err != nil {
		return data, err
	}

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}
