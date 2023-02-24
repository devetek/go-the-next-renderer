package model

// Shop Response Struct

type ResponseShopDetail struct {
	Shop     Shop      `json:"shop"`
	Products []Product `json:"products"`
}

// Product Response Struct
type ResponseProductList struct {
	Products Product `json:"products"`
}

type ResponseProductDetail struct {
	Product Product   `json:"product"`
	Related []Product `json:"related"`
}
