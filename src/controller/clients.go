package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/service"
)

type ClientController struct{}

func (cl ClientController) PostTransactions(c *gin.Context) {
	var request service.TransactionRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}
	service.PostTransaction(&request, c.Param("id"))

}

func (cl ClientController) GetBankStatement(c *gin.Context) {

}
