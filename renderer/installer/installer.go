package installer

import (
	"io"
	"log"
	"net/http"
	"os"
)

type TemplateSource struct {
	Endpoint string
}

func New() TemplateSource {
	var source TemplateSource
	source.Endpoint = "http://localhost:3000"

	source.Generate()

	return source
}

// TODO: get list of template from registered definition
func (source *TemplateSource) Generate() {
	var client = &http.Client{}

	request, err := http.NewRequest("GET", source.Endpoint, nil)
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("user-agent", "wpe-bot-aggregator-v0.0.1")

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}

	f, err := os.Create("template/home.html")
	if err != nil {
		log.Println(err)
	}

	defer f.Close()
	defer response.Body.Close()

	io.Copy(f, response.Body)
}
