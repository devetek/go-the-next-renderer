package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
}
