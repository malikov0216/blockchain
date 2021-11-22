package repository

import (
	"blockchain/pkg/model"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CurrencyPostgres struct {
	db *sqlx.DB
}

func NewCurrencyPostgres(db *sqlx.DB) *CurrencyPostgres {
	return &CurrencyPostgres{db}
}

func (r *CurrencyPostgres) Create(currency model.Currency, tx *sqlx.Tx) (int64, error) {
	var id int64

	query := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", currencyTable)
	row := tx.QueryRow(query, currency.Name)

	if err := row.Scan(&id); err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	return id, nil
}

func (r *CurrencyPostgres) GetCurrencies() ([]model.Currency, error) {
	var currencies []model.Currency
	query := fmt.Sprintf("SELECT id, name FROM %s where deleted_at ISNULL", currencyTable)
	err := r.db.Select(&currencies, query)
	return currencies, err
}

func (r *CurrencyPostgres) GetCurrencyByID(currencyID int64) (model.Currency, error) {
	currency := model.Currency{}
	query := fmt.Sprintf("SELECT id, name FROM %s where id = $1 and deleted_at ISNULL", currencyTable)
	err := r.db.Get(&currency, query, currencyID)
	return currency, err
}
