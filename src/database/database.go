package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

var dbPool *pgxpool.Pool

func Init() {
	var err error
	dbPool, err = pgxpool.New(context.Background(), os.Getenv("API_DATABASE_URI"))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
}

func GetDatabase() *pgxpool.Pool {
	return dbPool
}
