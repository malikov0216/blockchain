package repository

import (
	"blockchain/pkg/model"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (r *AuthPostgres) Create(user model.User, tx *sqlx.Tx) (int64, error) {
	var id int64

	query := fmt.Sprintf("INSERT INTO %s (email, password, first_name, second_name) values ($1, $2, $3, $4) RETURNING id", userTable)
	row := tx.QueryRow(query, user.Email, user.Password, user.FirstName, user.SecondName)

	if err := row.Scan(&id); err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id, password FROM %s WHERE email=$1 and deleted_at ISNULL", userTable)
	err := r.db.Get(&user, query, email)
	return user, err
}
