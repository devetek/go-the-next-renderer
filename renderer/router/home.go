package router

import (
	"log"
	"net/http"

	"github.com/devetek/go-the-next-renderer/template"
	"github.com/devetek/go-the-next-renderer/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := utils.Fetcher()
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
