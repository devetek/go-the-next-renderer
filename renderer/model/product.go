package model

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"category,omitempty"`
	Price       string `json:"price,omitempty"`
	PriceRaw    int    `json:"price_raw,omitempty"`
	Discount    string `json:"discount,omitempty"`
	Description string `json:"description,omitempty"`
}
