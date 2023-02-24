package template

import (
	"io"

	"github.com/devetek/go-the-next-renderer/model"
)

// Shop Page
type ShopDetailParams model.ResponseShopDetail

func Shop(w io.Writer, p ShopDetailParams) error {
	return parse("shop").Execute(w, p)
}
