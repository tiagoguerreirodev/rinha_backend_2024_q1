package service

import (
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/constant"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/database"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/model"
	"log"
	"time"
)

func PostTransaction(request *model.TransactionRequest, userId string) (*model.TransactionResponse, error) {
	client, err := database.GetUser(userId)
	if err != nil {
		log.Fatalf("Could not select user: %v\n", err)
	}

	var transactionValue int

	if request.Type == "d" {
		transactionValue = -request.Value
	} else {
		transactionValue = request.Value
	}

	updatedBalance := client.Balance + transactionValue

	if request.Type == "d" && updatedBalance < -client.Limit {
		return nil, constant.ErrBalanceExceedsLimit
	}

	if err = database.SaveTransaction(request, &userId); err != nil {
		return nil, err
	}

	if err = database.UpdateUserBalance(userId, &updatedBalance); err != nil {
		return nil, err
	}

	res := &model.TransactionResponse{
		Limite: client.Limit,
		Saldo:  updatedBalance,
	}

	return res, nil
}

func GetStatement(id string) *model.BankStatementResponse {
	client, clientErr := database.GetUser(id)
	if clientErr != nil {
		log.Fatalf("Could not select user: %v\n", clientErr)
	}

	transactions, transactionErr := database.GetBankStatement(&id)
	if transactionErr != nil {
		log.Fatalf("Could not retrieve bank statement: %v\n", transactionErr)
	}

	statement := &model.Statement{
		Total:  client.Balance,
		Data:   time.Now().String(),
		Limite: client.Limit,
	}

	return &model.BankStatementResponse{
		Saldo:      statement,
		Transacoes: transactions,
	}
}
