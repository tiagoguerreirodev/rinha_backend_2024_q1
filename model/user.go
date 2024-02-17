package model

type User struct {
	Id      string `json:"id"`
	Balance int    `json:"balance"`
	Limit   int    `json:"limit"`
}
