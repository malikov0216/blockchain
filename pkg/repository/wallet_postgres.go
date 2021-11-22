package repository

import (
	"blockchain/pkg/model"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type WalletPostgres struct {
	db *sqlx.DB
}

func NewWalletPostgres(db *sqlx.DB) *WalletPostgres {
	return &WalletPostgres{db}
}

func (r *WalletPostgres) GetWalletsByUserID(userID int64) ([]model.Wallet, error) {
	wallets := []model.Wallet{}
	query := fmt.Sprintf("SELECT id, address, currency_id, balance, user_id from %s where user_id = $1 and deleted_at ISNULL", walletTable)
	err := r.db.Select(&wallets, query, userID)
	return wallets, err
}

func (r *WalletPostgres) GetUserWalletByAddress(address string, userID int64, tx *sqlx.Tx) (*model.Wallet, error) {
	wallet := new(model.Wallet)
	query := fmt.Sprintf("SELECT id, address, currency_id, balance, user_id from %s where address = $1 and user_id = $2 and deleted_at ISNULL FOR UPDATE", walletTable)
	row := tx.QueryRow(query, address, userID)
	if err := row.Scan(wallet); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return wallet, nil
}

func (r *WalletPostgres) Create(wallet model.Wallet, tx *sqlx.Tx) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (address, currency_id, balance, user_id) values ($1, $2, $3, $4) RETURNING id", walletTable)
	row := tx.QueryRow(query, wallet.Address, wallet.CurrencyID, wallet.Balance, wallet.UserID)
	if err := row.Scan(&id); err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return id, nil
}

func (r *WalletPostgres) UpdateBalance(wallet *model.Wallet, tx *sqlx.Tx) error {
	query := fmt.Sprintf("UPDATE %s SET balance = $1, updated_at = $2 WHERE id = $3", walletTable)
	_, err := tx.Exec(query, wallet.Balance, time.Now(), wallet.ID)
	if err != nil {
		return err
	}
	return nil
}
