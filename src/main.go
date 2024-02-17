package main

import (
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/database"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/server"
)

func main() {
	database.Init()
	defer database.GetDatabase().Close()
	server.Init()
}
