package model

type Transaction struct {
	UserId      string `json:"userId"`
	Type        string `json:"type"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	Value       int16  `json:"value"`
}

type TransactionRequest struct {
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
	Value       int    `json:"valor"`
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
	Data   string `json:"data_extrato"`
	Total  int    `json:"total"`
	Limite int    `json:"limite"`
}
