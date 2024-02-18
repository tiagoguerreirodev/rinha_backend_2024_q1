package service

import (
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/cache"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/constant"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/database"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/model"
	"log"
	"math"
	"time"
)

var c *cache.AllCache

func init() {
	c = cache.GetCache()
}

func PostTransaction(request *model.TransactionRequest, userId string) (*model.TransactionResponse, error) {
	client := c.ReadClient(userId)

	var transactionValue int

	if request.Type == "d" {
		transactionValue = -request.Value
	} else {
		transactionValue = request.Value
	}

	updatedBalance := client.Balance + transactionValue

	if request.Type == "d" && math.Abs(float64(updatedBalance)) > float64(client.Limit) {
		return nil, constant.ErrBalanceExceedsLimit
	}

	if dbErr := database.SaveTransaction(request, &userId); dbErr != nil {
		log.Fatalf("Could not insert transaction: %v\n", dbErr)
	}

	res := &model.TransactionResponse{
		Limite: client.Limit,
		Saldo:  updatedBalance,
	}

	c.UpdateClient(userId, &model.User{
		Balance: res.Saldo,
		Limit:   res.Limite,
	})

	return res, nil
}

func GetStatement(id string) *model.BankStatementResponse {
	client := c.ReadClient(id)

	transactions, err := database.GetBankStatement(&id)
	if err != nil {
		log.Fatalf("Could not retrieve bank statement: %v\n", err)
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
