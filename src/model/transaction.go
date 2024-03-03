package model

type Transaction struct {
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
	CreatedAt   string `json:"realizada_em"`
	Value       int16  `json:"valor"`
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
	Saldo      *Statement    `json:"saldo"`
	Transacoes []Transaction `json:"ultimas_transacoes"`
}

type Statement struct {
	Data   string `json:"data_extrato"`
	Total  int    `json:"total"`
	Limite int    `json:"limite"`
}

func (t *TransactionRequest) GetTransactionValue() int {
	switch t.Type {
	case "c":
		return t.Value
	case "d":
		return -t.Value
	default:
		return 0
	}
}
