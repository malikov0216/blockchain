package handler

import (
	"blockchain/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createCurrency(c *gin.Context) {
	tx, err := h.db.Beginx()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer tx.Commit()

	var currency model.Currency
	if err := c.BindJSON(&currency); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := currency.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	currencyID, err := h.services.Currency.Create(currency, tx)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		tx.Rollback()
		return
	}

	newSuccessResponse(c, "created new currency", currencyID)
}
func (h *Handler) getCurrencies(c *gin.Context) {
	currencies, err := h.services.GetCurrencies()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error while getting currencies data")
	}
	newSuccessResponse(c, "received currency list", currencies)
}
