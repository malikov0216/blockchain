package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserWallets(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	wallets, err := h.services.GetUserWallets(userID.(int64))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newSuccessResponse(c, "received users wallet", wallets)
}
