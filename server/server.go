package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/controller"
	"log"
	"os"
)

func Init() {
	router := gin.Default()

	clientController := new(controller.ClientController)

	router.POST("/clientes/:id/transacoes", clientController.PostTransactions)
	router.GET("/clientes/:id/extrato", clientController.GetBankStatement)

	if err := router.Run(os.Getenv("API_URI")); err != nil {
		log.Fatalf("Failed to run application: %v\n", err)
	}
}
