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
	loop    = parse("loop.html")
)

// Object Page
type HomeParams utils.ResponseUser

func Home(w io.Writer, p HomeParams) error {
	return home.Execute(w, p)
}

// Loop Page
type LoopParams utils.ResponseTopics

func Loop(w io.Writer, p LoopParams) error {
	return loop.Execute(w, p)
}

// Example Page
type ExampleParams struct {
	Title   string
	Message string
}

func Example(w io.Writer, p ExampleParams) error {
	return example.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.ParseFS(files, file))
}
