package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Wallet struct {
	ID         int64      `json:"id,omitempty" db:"id"`
	Address    string     `json:"address,omitempty"  db:"address"`
	CurrencyID int64      `json:"currency_id,omitempty" db:"currency_id" `
	Balance    float64    `json:"balance,omitempty" db:"balance" `
	UserID     int64      `json:"user_id,omitempty" db:"user_id" `
	CreatedAt  *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (w Wallet) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.UserID, validation.Required),
		validation.Field(&w.CurrencyID, validation.Required),
	)
}
