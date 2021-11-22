package handler

import (
	"blockchain/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	tx, err := h.db.Beginx()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer tx.Commit()

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := user.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := h.services.Authorization.Create(user, tx)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		tx.Rollback()
		return
	}

	currencies, err := h.services.Currency.GetCurrencies()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		tx.Rollback()
		return
	}

	for i := range currencies {
		wallet := model.Wallet{
			CurrencyID: currencies[i].ID,
			UserID:     userID,
			Balance:    100,
		}

		if err := wallet.Validate(); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			tx.Rollback()
			return
		}

		_, err := h.services.Wallet.Create(wallet, tx)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			tx.Rollback()
			return
		}
	}

	newSuccessResponse(c, "registered", userID)
}

func (h *Handler) signIn(c *gin.Context) {
	var user model.UserCredential

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(user.Email, user.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newSuccessResponse(c, "authorized", token)
}
