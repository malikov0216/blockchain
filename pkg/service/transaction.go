package service

import (
	"blockchain/pkg/model"
	"blockchain/pkg/repository"
	"database/sql"
	"fmt"

	"github.com/dariubs/percent"
	"github.com/jmoiron/sqlx"
)

type TransactionService struct {
	repo         repository.Transaction
	walletRepo   repository.Wallet
	currencyRepo repository.Currency
}

func NewTransactionService(repo repository.Transaction, walletRepo repository.Wallet, currencyRepo repository.Currency) *TransactionService {
	return &TransactionService{repo, walletRepo, currencyRepo}
}

func (s *TransactionService) GetUserTransactions(userID int64) ([]model.Transaction, error) {
	return s.repo.GetTransactionsBy(userID)
}

func (s *TransactionService) TransferFund(transaction model.Transaction, tx *sqlx.Tx) (int64, error) {
	walletFrom, err := s.walletRepo.GetUserWalletByAddress(transaction.AddressFrom, transaction.UserID, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("wallet creds from where you want transfer funds is not right")
		}
		return 0, err
	}

	walletTo, err := s.walletRepo.GetUserWalletByAddress(transaction.AddressTo, transaction.UserID, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("wallet creds where you want to transfer funds is not right")
		}
		return 0, err
	}

	commissionAmmount := percent.Percent(int(transaction.Commission), int(transaction.Sum))

	if walletFrom.Balance < float64(transaction.Sum)+commissionAmmount {
		return 0, fmt.Errorf("you don't have enough funds")
	}

	if walletFrom.CurrencyID != walletTo.CurrencyID {
		return 0, fmt.Errorf("the currency to which you want to transfer is not suitable for transfer")
	}

	currency, err := s.currencyRepo.GetCurrencyByID(walletFrom.CurrencyID)
	if err != nil {
		return 0, err
	}
	transaction.CurrencyName = currency.Name

	walletFrom.Balance -= commissionAmmount + float64(transaction.Sum)
	err = s.walletRepo.UpdateBalance(walletFrom, tx)
	if err != nil {
		return 0, err
	}

	walletTo.Balance += float64(transaction.Sum)
	err = s.walletRepo.UpdateBalance(walletTo, tx)
	if err != nil {
		return 0, err
	}

	id, err := s.repo.Create(transaction, tx)
	if err != nil {
		return 0, err
	}

	return id, nil
}
