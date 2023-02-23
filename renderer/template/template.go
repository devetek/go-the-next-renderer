package template

import (
	"embed"
	"io"
	"text/template"

	"github.com/devetek/go-the-next-renderer/utils"
)

//go:embed *
var files embed.FS

var (
	home    = parse("home.html")
	example = parse("example.html")
)

type HomeParams utils.Response

func Home(w io.Writer, p HomeParams) error {
	return home.Execute(w, p)
}

type ExampleParams struct {
	Title   string
	Message string
}

func Dashboard(w io.Writer, p ExampleParams) error {
	return example.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.ParseFS(files, file))
}
