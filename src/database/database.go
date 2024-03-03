package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/constant"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/model"
	"log"
	"os"
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

func UpdateUserBalance(
	request *model.TransactionRequest,
	userId string,
) (*model.TransactionResponse, error) {

	var updatedBalance, limite int

	if err := dbPool.QueryRow(
		context.Background(),
		"select * from updateUserBalance($1,$2)",
		userId,
		request.GetTransactionValue(),
	).Scan(&updatedBalance, &limite); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23514" && pgErr.ConstraintName == "balance_within_limit" {
				return nil, constant.ErrBalanceExceedsLimit
			}
			if pgErr.Message == "user not found on postgres" {
				return nil, constant.ErrUserNotFound
			}
		}
		return nil, err
	}

	return &model.TransactionResponse{
		Limite: limite,
		Saldo:  updatedBalance,
	}, nil
}

func SaveTransaction(
	request *model.TransactionRequest,
	userId string,
) {
	_, _ = dbPool.Exec(
		context.Background(),
		"insert into transacoes (user_id, valor, tipo, descricao) values ($1,$2,$3,$4)",
		userId,
		request.Value,
		request.Type,
		request.Description,
	)
	return
}

func GetUser(
	id string,
) (*model.User, error) {
	var saldo, limite int

	err := dbPool.QueryRow(
		context.Background(),
		"select saldo, limite from clientes WHERE id = $1",
		id,
	).Scan(&saldo, &limite)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, constant.ErrUserNotFound
		}
		return nil, err
	}

	return &model.User{
		Balance: saldo,
		Limit:   limite,
	}, nil
}

func GetBankStatement(
	userId string,
) ([]model.Transaction, error) {

	rows, err := dbPool.Query(
		context.Background(),
		"select valor, tipo, descricao, created_at from transacoes where user_id = $1 order by id desc limit 10",
		userId,
	)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var transactions []model.Transaction
	for rows.Next() {
		var valor int16
		var tipo, descricao string
		var dataCriacao pgtype.Timestamp
		err = rows.Scan(&valor, &tipo, &descricao, &dataCriacao)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, model.Transaction{
			Value:       valor,
			Type:        tipo,
			Description: descricao,
			CreatedAt:   dataCriacao.Time.String(),
		})
	}
	return transactions, nil
}
