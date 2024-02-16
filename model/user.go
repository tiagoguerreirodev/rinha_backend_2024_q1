package model

type User struct {
	Id      int8 `json:"id"`
	Balance int  `json:"balance"`
	Limit   int  `json:"limit"`
}
