package handler

import (
	"blockchain/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

const commission = 50

func (h *Handler) getUserTransactions(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	transactions, err := h.services.GetUserTransactions(userID.(int64))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if len(transactions) == 0 {
		newErrorResponse(c, http.StatusForbidden, "you dont have transfered funds yes")
		return
	}
	newSuccessResponse(c, "received users wallet", transactions)
}

func (h *Handler) transferFunds(c *gin.Context) {
	tx, err := h.db.Beginx()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer tx.Commit()
	var transferFundBody model.TransferFundBody
	var transaction model.Transaction

	userID, _ := c.Get(userCtx)

	if err := c.BindJSON(&transferFundBody); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := transferFundBody.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	transaction.TransferFundBody = transferFundBody
	transaction.UserID = userID.(int64)
	transaction.Commission = commission

	id, err := h.services.TransferFund(transaction, tx)
	if err != nil {
		tx.Rollback()
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newSuccessResponse(c, "transfered funds", id)
}
