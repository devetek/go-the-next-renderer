package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devetek/go-the-next-renderer/model"
	"github.com/devetek/go-the-next-renderer/template"
	"github.com/devetek/go-the-next-renderer/utils"
	"github.com/gorilla/mux"
)

// get link API
func getLinkProductDetailAPI(target string) string {
	apiMap := make(map[string]string)
	apiMap["a"] = "https://mocki.io/v1/75805603-5327-4a4e-af47-af8307988c35"
	apiMap["b"] = "https://mocki.io/v1/dc40078b-3d11-426c-9f61-c41c2e2428de"
	apiMap["default"] = "https://mocki.io/v1/dcb503d6-671a-4ecd-b5ee-443dbd8806db"

	selectedData := apiMap[target]

	if selectedData == "" {
		return apiMap["default"]
	}

	return selectedData
}

// Product dandler
func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var data model.ResponseProductDetail

	response, err := utils.Fetcher(getLinkProductDetailAPI(vars["name"]))
	if err != nil {
		log.Println(err)
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}

	p := template.ProductDetailParams{
		Product: data.Product,
		Related: data.Related,
	}
	template.ProductDetail(w, p)
}
