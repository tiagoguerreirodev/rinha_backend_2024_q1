package main

import (
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/database"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/server"
)

func main() {
	defer database.Pool().Close()
	server.Init()
}
