package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/constant"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/model"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/service"
	"net/http"
	"strconv"
)

type ClientController struct{}

func (cl ClientController) PostTransactions(c *gin.Context) {
	id := c.Param("id")

	if isUserIdInvalid(id) {
		c.Status(http.StatusNotFound)
		return
	}
	var request model.TransactionRequest

	if err := c.BindJSON(&request); err != nil {
		return
	}

	if len(request.Description) == 0 || len(request.Description) > 10 {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	if request.Type != "c" && request.Type != "d" {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	response, err := service.PostTransaction(&request, id)

	if err != nil {
		if errors.Is(err, constant.ErrBalanceExceedsLimit) {
			c.Status(http.StatusUnprocessableEntity)
			return
		}
	}

	c.JSON(http.StatusOK, response)

}

func (cl ClientController) GetBankStatement(c *gin.Context) {
	id := c.Param("id")

	if isUserIdInvalid(id) {
		c.Status(http.StatusNotFound)
		return
	}

	response := service.GetStatement(id)
	c.JSON(http.StatusOK, response)
}

func isUserIdInvalid(id string) bool {
	integerId, _ := strconv.Atoi(id)
	return integerId < 1 || integerId > 5
}
