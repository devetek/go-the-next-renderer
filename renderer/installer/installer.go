package installer

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type TemplateSource struct {
	Module   string `json:"module"`
	Endpoint string `json:"endpoint"`
}

func New() error {
	var contracts []TemplateSource
	contractFile, err := os.Open("installer/contract.json")
	// Let it down, we don't want has invalid contract file
	if err != nil {
		return err
	}
	// defer the closing of our contractFile so that we can parse it later on
	defer contractFile.Close()

	jsonBytes, _ := io.ReadAll(contractFile)

	json.Unmarshal(jsonBytes, &contracts)

	// Generate contract templates
	for _, contract := range contracts {
		log.Println("contract: ", contract.Module)
		contract.Generate()
	}

	return err
}

// TODO: get list of template from registered definition
func (source *TemplateSource) Generate() {
	var client = &http.Client{}

	request, err := http.NewRequest("GET", source.Endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("user-agent", "wpe-bot-aggregator-v0.0.1")

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("template/" + source.Module + ".html")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	defer response.Body.Close()

	io.Copy(f, response.Body)
}
