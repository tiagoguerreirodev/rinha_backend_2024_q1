package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/model"
	"log"
	"os"
	"time"
)

var dbPool *pgxpool.Pool

func init() {
	var err error
	dbPool, err = pgxpool.New(context.Background(), os.Getenv("API_DATABASE_URI"))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
}

func Pool() *pgxpool.Pool {
	return dbPool
}

func SaveTransaction(request *model.TransactionRequest, userId *string) error {
	_, err := dbPool.Exec(
		context.Background(),
		"insert into transacoes values ($1,$2,$3,$4,$5)",
		userId,
		request.Value,
		request.Type,
		request.Description,
		time.Now(),
	)
	return err
}

func GetUser(id string) (*model.User, error) {
	var saldo, limite int

	err := dbPool.QueryRow(
		context.Background(),
		"select saldo, limite from clientes WHERE id = $1",
		id,
	).Scan(&saldo, &limite)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Balance: saldo,
		Limit:   limite,
	}, nil
}

func UpdateUserBalance(id string, newBalance *int) error {
	_, err := dbPool.Exec(
		context.Background(),
		"update clientes set saldo = $1 where id = $2",
		newBalance,
		id,
	)
	return err
}

func GetBankStatement(userId *string) ([]*model.Transaction, error) {

	rows, err := dbPool.Query(
		context.Background(),
		"select valor, tipo, descricao, created_at from transacoes where user_id = $1 order by created_at desc limit 10",
		userId,
	)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	transactions := make([]*model.Transaction, 0, 10)
	for rows.Next() {
		var valor int16
		var tipo, descricao string
		var dataCriacao pgtype.Timestamp
		errScan := rows.Scan(&valor, &tipo, &descricao, &dataCriacao)
		if errScan != nil {
			return nil, errScan
		}
		transactions = append(transactions, &model.Transaction{
			Value:       valor,
			Type:        tipo,
			Description: descricao,
			CreatedAt:   dataCriacao.Time.String(),
		})
	}

	return transactions, nil
}
