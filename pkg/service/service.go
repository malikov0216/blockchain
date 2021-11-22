package service

import (
	"blockchain/pkg/model"
	"blockchain/pkg/repository"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Create(user model.User, tx *sqlx.Tx) (int64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int64, error)
}

type Wallet interface {
	Create(wallet model.Wallet, tx *sqlx.Tx) (int64, error)
	GetUserWallets(userID int64) ([]model.Wallet, error)
}

type Transaction interface {
	GetUserTransactions(userID int64) ([]model.Transaction, error)
	TransferFund(transaction model.Transaction, tx *sqlx.Tx) (int64, error)
}

type Currency interface {
	GetCurrencies() ([]model.Currency, error)
	Create(currency model.Currency, tx *sqlx.Tx) (int64, error)
}

type Service struct {
	Authorization
	Wallet
	Currency
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Wallet:        NewWalletService(repos.Wallet),
		Currency:      NewCurrencyService(repos.Currency),
		Transaction:   NewTransactionService(repos.Transaction, repos.Wallet, repos.Currency),
	}
}
