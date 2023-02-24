package utils

import (
	"log"
	"net/http"
	"time"
)

type ResponseUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
}

func Fetcher(url string) (*http.Response, error) {
	var start = time.Now()
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

	elapsed := time.Since(start)
	log.Printf("Fetcher took %s", elapsed)

	return response, nil
}
