package service

import (
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/database"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/model"
	"time"
)

func PostTransaction(request *model.TransactionRequest, userId string) (*model.TransactionResponse, error) {

	database.SaveTransaction(request, userId)

	balance, err := database.UpdateUserBalance(request, userId)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func GetStatement(id string) (*model.BankStatementResponse, error) {

	client, err := database.GetUser(id)
	if err != nil {
		return nil, err
	}

	transactions, err := database.GetBankStatement(id)
	if err != nil {
		return nil, err
	}

	statement := &model.Statement{
		Total:  client.Balance,
		Data:   time.Now().String(),
		Limite: client.Limit,
	}

	return &model.BankStatementResponse{
		Saldo:      statement,
		Transacoes: transactions,
	}, nil
}
