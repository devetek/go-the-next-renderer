package template

import "io"

// Example Page
type ExampleParams struct {
	Title   string
	Message string
}

func Example(w io.Writer, p ExampleParams) error {
	return parse("example").Execute(w, p)
}
