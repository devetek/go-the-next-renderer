package template

import (
	"io"

	"github.com/devetek/go-the-next-renderer/model"
)

// Product Detail Page
type ProductDetailParams model.ResponseProductDetail

func ProductDetail(w io.Writer, p ProductDetailParams) error {
	return parse("product-detail").Execute(w, p)
}
