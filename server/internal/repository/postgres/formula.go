package postgres

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"server/internal/entities"
)

func DBFormulaGetByID(db *sqlx.DB, id int64) (*entities.Formula, error) {
	formula := entities.Formula{}
	query := `SELECT id, title, value, user_id FROM formula WHERE id = $1`
	err := db.Get(&formula, query, id)
	if err != nil {
		return &entities.Formula{}, nil
	}

	return &formula, nil
}

func DBFormulaGetByUserID(db *sqlx.DB, userID int64) (*[]entities.Formula, error) {
	formulas := []entities.Formula{}
	query := `SELECT id, title, value, user_id FROM formula WHERE user_id = $1`

	err := db.Select(&formulas, query, userID)
	if err != nil {
		return nil, err
	}
	return &formulas, nil
}

func DBFormulaCreate(db *sqlx.DB, formula *entities.Formula) (*entities.Formula, error) {
	query := `INSERT INTO formula (title, value, user_id)
	VALUES (:title, :value, :user_id) RETURNING id`

	stmt, err := db.PrepareNamed(query)
	if stmt == nil {
		return nil, errors.New("error preparing statement")
	}
	err = stmt.Get(&formula.ID, *formula)
	if err != nil {
		return nil, err
	}

	return formula, nil
}

func DBFormulaUpdate(db *sqlx.DB, formula *entities.UpdateFormulaRequest) error {
	query := `UPDATE formula SET title = $1, value = $2 WHERE id = $3`

	_, err := db.Exec(query, formula.Title, formula.Value, formula.ID)
	if err != nil {
		return err
	}
	return nil
}

func DBFormulaDelete(db *sqlx.DB, formulaID int64) error {
	query := `DELETE FROM formula WHERE id = $1`

	_, err := db.Exec(query, formulaID)
	if err != nil {
		return err
	}
	return nil
}

func DBFormulaHistoryGet(db *sqlx.DB, userID int64, pageNumber int64) (*[]entities.Formula, error) {
	formulas := []entities.Formula{}
	query := `
		SELECT id, title, value, user_id FROM formula WHERE user_id = $1
		ORDER BY id DESC
		LIMIT 20
		OFFSET $2
	`

	err := db.Select(&formulas, query, userID, pageNumber-1)
	if err != nil {
		return nil, err
	}
	return &formulas, nil
}
