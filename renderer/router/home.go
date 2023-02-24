package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devetek/go-the-next-renderer/template"
	"github.com/devetek/go-the-next-renderer/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var data utils.ResponseUser
	response, err := utils.Fetcher("https://mocki.io/v1/4459ceb7-791b-4948-87e1-db598b805587")
	if err != nil {
		log.Println(err)
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}

	p := template.HomeParams{
		ID:       data.ID,
		Username: data.Username,
		Fullname: data.Fullname,
		Address:  data.Address,
	}
	template.Home(w, p)
}
