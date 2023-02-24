package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devetek/go-the-next-renderer/template"
	"github.com/devetek/go-the-next-renderer/utils"
)

func LoopHandler(w http.ResponseWriter, r *http.Request) {
	var data utils.ResponseTopics
	response, err := utils.Fetcher("https://mocki.io/v1/15b1a310-fbe6-453a-9fcb-57ec3e240cb1")
	if err != nil {
		log.Println(err)
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}

	p := template.LoopParams{
		Topics: data.Topics,
	}
	template.Loop(w, p)
}
