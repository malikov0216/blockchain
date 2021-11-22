package service

import (
	"blockchain/pkg/model"
	"blockchain/pkg/repository"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type WalletService struct {
	repo repository.Wallet
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{repo}
}

func (s *WalletService) GetUserWallets(userID int64) ([]model.Wallet, error) {
	return s.repo.GetWalletsByUserID(userID)
}

func (s *WalletService) Create(wallet model.Wallet, tx *sqlx.Tx) (int64, error) {
	sha256Address, err := generateSHA256()
	if err != nil {
		return 0, err
	}
	wallet.Address = sha256Address
	return s.repo.Create(wallet, tx)
}

func generateSHA256() (string, error) {
	data := make([]byte, 10)
	if _, err := rand.Read(data); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum256(data)), nil
}
