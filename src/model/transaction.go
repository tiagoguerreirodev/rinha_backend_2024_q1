package model

type Transaction struct {
	UserId      string `json:"userId"`
	Value       int16  `json:"value"`
	Type        string `json:"type"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

type TransactionRequest struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

type TransactionResponse struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}

type BankStatementResponse struct {
	Saldo      *Statement     `json:"saldo"`
	Transacoes []*Transaction `json:"ultimas_transacoes"`
}

type Statement struct {
	Total  int    `json:"total"`
	Data   string `json:"data_extrato"`
	Limite int    `json:"limite"`
}
