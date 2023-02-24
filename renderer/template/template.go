package template

import (
	"embed"
	"text/template"
)

//go:embed *
var files embed.FS

func parse(file string) *template.Template {
	return template.Must(
		template.ParseFS(files, file+".html"))
}
