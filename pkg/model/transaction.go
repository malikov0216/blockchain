package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Transaction struct {
	ID           int64  `json:"id,omitempty" db:"id"`
	Commission   int64  `json:"commission,omitempty" db:"commission" `
	UserID       int64  `json:"user_id,omitempty"  db:"user_id"`
	CurrencyName string `json:"currency_name,omitempty" db:"currency_name"`
	TransferFundBody
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type TransferFundBody struct {
	AddressFrom string `json:"address_from,omitempty" db:"address_from"`
	AddressTo   string `json:"address_to,omitempty" db:"address_to"`
	Sum         int64  `json:"sum,omitempty" db:"sum"`
}

func (t TransferFundBody) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.AddressFrom, validation.Required),
		validation.Field(&t.AddressTo, validation.Required),
		validation.Field(&t.Sum, validation.Required),
	)
}
