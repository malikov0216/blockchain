package handler

import (
	"blockchain/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	services *service.Service
	db       *sqlx.DB
}

func NewHandler(service *service.Service, db *sqlx.DB) *Handler {
	return &Handler{services: service, db: db}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIndentity)
	{
		api.GET("/currencies", h.getCurrencies)
		api.POST("/currency", h.createCurrency)
		api.GET("/wallets", h.getUserWallets)
		api.GET("/transactions", h.getUserTransactions)
		api.POST("/transaction", h.transferFunds)
	}
	return router
}
