package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

func Init() {
	conn, err := pgxpool.New(context.Background(), os.Getenv("API_DATABASE_URI"))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()
}
