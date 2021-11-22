package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	ID int64 `json:"id,omitempty" db:"id"`
	UserCredential
	FirstName  string     `json:"first_name,omitempty" db:"first_name"`
	SecondName string     `json:"second_name,omitempty" db:"second_name"`
	CreatedAt  *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type UserCredential struct {
	Email    string `json:"email,omitempty" db:"email"`
	Password string `json:"password,omitempty" db:"password"`
}

func (u User) Validate() error {
	isExistingEmail := validation.NewStringRule(govalidator.IsExistingEmail, "must be real domain email")
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, isExistingEmail),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 50)),
		validation.Field(&u.FirstName, validation.Required, is.UTFLetterNumeric),
		validation.Field(&u.SecondName, validation.Required, is.UTFLetterNumeric),
	)
}
