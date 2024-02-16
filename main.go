package main

import (
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/database"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/server"
)

func main() {

	database.Init()
	server.Init()
}
