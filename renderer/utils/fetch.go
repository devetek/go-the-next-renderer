package utils

import (
	"net/http"
)

type ResponseUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
}

func Fetcher(url string) (*http.Response, error) {
	var client = &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// add body close in your invoker
	// defer response.Body.Close()

	return response, nil
}
