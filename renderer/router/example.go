package router

import (
	"net/http"

	"github.com/devetek/go-the-next-renderer/template"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	p := template.ExampleParams{
		Title:   "Example Page",
		Message: "Hello from example",
	}
	template.Example(w, p)
}
