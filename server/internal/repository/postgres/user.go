package postgres

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"server/internal/entities"
)

func DBUserGetById(db *sqlx.DB, id int64) (*entities.User, error) {
	user := entities.User{}
	query := `SELECT FROM users (id, login, email, password) 
	WHERE id=$1`
	err := db.Get(&user, query, id)
	if err != nil {
		return &entities.User{}, nil
	}

	return &user, nil
}

func DBUserGetByLogin(db *sqlx.DB, login string) (*entities.User, error) {
	user := entities.User{}
	query := `SELECT id, email, login, password FROM users WHERE login = $1`
	err := db.Get(&user, query, login)
	if err != nil {
		return &entities.User{}, nil
	}

	return &user, nil

}

func DBUserExists(db *sqlx.DB, login string, email string) (bool, error) {
	exists := 0
	query := `SELECT 1 FROM users WHERE login=$1 OR email = $2 LIMIT 1`

	err := db.QueryRow(query, login, email).Scan(&exists)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

func DBUserCreate(db *sqlx.DB, user *entities.User) (*entities.User, error) {
	query := `INSERT INTO users (login, email, password)
	VALUES (:login, :email, :password) RETURNING id`

	stmt, err := db.PrepareNamed(query)
	if stmt == nil {
		return nil, errors.New("error preparing statement")
	}
	err = stmt.Get(&user.ID, *user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
