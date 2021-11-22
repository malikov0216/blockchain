package repository

import (
	"blockchain/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Create(user model.User, tx *sqlx.Tx) (int64, error)
	GetUser(username string) (model.User, error)
}

type Wallet interface {
	Create(wallet model.Wallet, tx *sqlx.Tx) (int64, error)
	GetWalletsByUserID(userID int64) ([]model.Wallet, error)
	GetUserWalletByAddress(address string, userID int64, tx *sqlx.Tx) (*model.Wallet, error)
	UpdateBalance(wallet *model.Wallet, tx *sqlx.Tx) error
}

type Transaction interface {
	GetTransactionsBy(userID int64) ([]model.Transaction, error)
	Create(transaction model.Transaction, tx *sqlx.Tx) (int64, error)
}

type Currency interface {
	Create(currency model.Currency, tx *sqlx.Tx) (int64, error)
	GetCurrencies() ([]model.Currency, error)
	GetCurrencyByID(currencyID int64) (model.Currency, error)
}

type Repository struct {
	Authorization
	Wallet
	Transaction
	Currency
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Wallet:        NewWalletPostgres(db),
		Transaction:   NewTransactionPostgres(db),
		Currency:      NewCurrencyPostgres(db),
	}
}
