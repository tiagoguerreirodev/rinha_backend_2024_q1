package service

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/database"
	"time"
)

type TransactionRequest struct {
	Value       int16  `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

var db *pgxpool.Pool

func PostTransaction(request *TransactionRequest, userId string) {
	db = database.GetDatabase()
	db.QueryRow(
		context.Background(),
		"insert into transaction values ($1, $2, $3, $4, $5)",
		userId,
		request.Value,
		request.Type,
		request.Description,
		time.Now(),
	)
}
