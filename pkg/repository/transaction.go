package repository

import (
	"blockchain/pkg/model"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db}
}

func (r *TransactionPostgres) GetTransactionsBy(userID int64) ([]model.Transaction, error) {
	transactions := []model.Transaction{}
	query := fmt.Sprintf("SELECT id, sum, commission, user_id, currency_name, address_from, address_to, created_at from %s where user_id = $1 and deleted_at ISNULL", transactionTable)
	err := r.db.Select(&transactions, query, userID)
	return transactions, err
}

func (r *TransactionPostgres) Create(transaction model.Transaction, tx *sqlx.Tx) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (sum, commission, user_id, currency_name, address_from, address_to) values ($1, $2, $3, $4, $5, $6) RETURNING id", transactionTable)
	row := tx.QueryRow(query, transaction.Sum, transaction.Commission, transaction.UserID, transaction.CurrencyName, transaction.AddressFrom, transaction.AddressTo)
	if err := row.Scan(&id); err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return id, nil
}
