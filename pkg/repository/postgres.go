package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	userTable        = "users"
	walletTable      = "wallets"
	transactionTable = "transactions"
	currencyTable    = "currencies"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s?sslmode=disable", cfg.DBName, cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	return sqlx.Connect("postgres", url)
}
