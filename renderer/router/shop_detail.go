package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/devetek/go-the-next-renderer/model"
	"github.com/devetek/go-the-next-renderer/template"
	"github.com/devetek/go-the-next-renderer/utils"
	"github.com/gorilla/mux"
)

// get link API
func getLinkShopDetailAPI(target string) string {
	apiMap := make(map[string]string)
	apiMap["a"] = "https://mocki.io/v1/61674e3c-675f-4cfe-a38d-cf676bffbda1"
	apiMap["b"] = "https://mocki.io/v1/07ac5dc4-ac87-4bf7-8053-f61de366d136"
	apiMap["default"] = "https://mocki.io/v1/d1f33761-e36c-410f-8cfb-9014caf6f3a1"

	selectedData := apiMap[target]

	if selectedData == "" {
		return apiMap["default"]
	}

	return selectedData
}

func ShopHandler(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var data model.ResponseShopDetail
	response, err := utils.Fetcher(getLinkShopDetailAPI(vars["name"]))
	if err != nil {
		log.Println(err)
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("data:", data.Shop)

	p := template.ShopDetailParams{
		Shop:     data.Shop,
		Products: data.Products,
	}
	template.Shop(w, p)
}
