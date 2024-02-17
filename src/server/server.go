package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/controller"
	"log"
)

func Init() {
	router := gin.Default()

	clientController := new(controller.ClientController)

	router.POST("/clientes/:id/transacoes", clientController.PostTransactions)
	router.GET("/clientes/:id/extrato", clientController.GetBankStatement)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run application: %v\n", err)
	}
}
