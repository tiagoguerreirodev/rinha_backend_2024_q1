package model

type Transaction struct {
	UserId      int8   `json:"userId"`
	Value       int16  `json:"value"`
	Type        string `json:"type"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}
