package main

import (
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/database"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/server"
)

func main() {
	database.Init()
	defer database.GetDatabase().Close()
	server.Init()
}
