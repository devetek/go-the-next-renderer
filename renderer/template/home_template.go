package template

import (
	"io"

	"github.com/devetek/go-the-next-renderer/model"
)

// Home Page
type HomeParams model.User

func Home(w io.Writer, p HomeParams) error {
	return parse("home").Execute(w, p)
}
