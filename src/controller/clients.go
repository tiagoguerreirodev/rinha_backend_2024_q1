package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/model"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/service"
	"log"
	"net/http"
	"strconv"
)

type ClientController struct{}

func (cl ClientController) PostTransactions(c *gin.Context) {
	id := c.Param("id")
	if isUserIdInvalid(id) {
		c.Status(http.StatusNotFound)
	}

	var request model.TransactionRequest
	if err := c.BindJSON(&request); err != nil {
		log.Fatalf("Could not bind JSON request body: %v\n", err)
	}

	response, err := service.PostTransaction(&request, id)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
	}

	c.JSON(http.StatusOK, response)

}

func (cl ClientController) GetBankStatement(c *gin.Context) {
	id := c.Param("id")
	if isUserIdInvalid(id) {
		c.Status(http.StatusNotFound)
	}

	response := service.GetStatement(id)
	c.JSON(http.StatusOK, response)
}

func isUserIdInvalid(id string) bool {
	integerId, _ := strconv.Atoi(id)
	return integerId < 1 || integerId > 5
}
