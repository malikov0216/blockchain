package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Currency struct {
	ID        int64      `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (c Currency) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, is.UpperCase),
	)
}
