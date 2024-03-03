package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/constant"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/model"
	"github.com/tiagoguerreirodev/rinha_backend_2024_q1/src/service"
	"net/http"
)

type ClientController struct{}

func (cl ClientController) PostTransactions(c *gin.Context) {

	var request model.TransactionRequest

	if err := c.BindJSON(&request); err != nil {
		return
	}

	if !isRequestValid(&request) {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	response, err := service.PostTransaction(&request, c.Param("id"))

	if err != nil {
		if errors.Is(err, constant.ErrBalanceExceedsLimit) {
			c.Status(http.StatusUnprocessableEntity)
		}
		if errors.Is(err, constant.ErrUserNotFound) {
			c.Status(http.StatusNotFound)
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

func (cl ClientController) GetBankStatement(c *gin.Context) {

	res, err := service.GetStatement(c.Param("id"))

	if err != nil {
		if errors.Is(err, constant.ErrUserNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
	}

	c.JSON(http.StatusOK, res)
}

func isRequestValid(request *model.TransactionRequest) bool {
	return request.Value > 0 &&
		(len(request.Description) > 0 && len(request.Description) <= 10) &&
		(request.Type == "c" || request.Type == "d")
}
