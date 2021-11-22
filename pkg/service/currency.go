package service

import (
	"blockchain/pkg/model"
	"blockchain/pkg/repository"

	"github.com/jmoiron/sqlx"
)

type CurrencyService struct {
	repo repository.Currency
}

func NewCurrencyService(repo repository.Currency) *CurrencyService {
	return &CurrencyService{repo}
}

func (s *CurrencyService) Create(currency model.Currency, tx *sqlx.Tx) (int64, error) {
	return s.repo.Create(currency, tx)
}

func (s *CurrencyService) GetCurrencies() ([]model.Currency, error) {
	return s.repo.GetCurrencies()
}
