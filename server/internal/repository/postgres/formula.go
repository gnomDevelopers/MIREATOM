package postgres

import (
	"database/sql"
	"errors"
	"server/internal/entities"
	"server/util"

	"github.com/jmoiron/sqlx"
)

// DBFormulaExists проверка существования формулы в бд
func DBFormulaExists(db *sqlx.DB, email string) (bool, error) {
	exists := 0
	query := `SELECT 1 FROM formula WHERE value = $1 LIMIT 1`

	err := db.QueryRow(query, email).Scan(&exists)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

// DBFormulaGetByID получение формулы по айди
func DBFormulaGetByID(db *sqlx.DB, id int64) (*entities.Formula, error) {
	formula := entities.Formula{}
	query := `SELECT id, title, value, user_id FROM formula WHERE id = $1`
	err := db.Get(&formula, query, id)
	if err != nil {
		return nil, err
	}

	return &formula, nil
}

// DBFormulaGetByUserID получение всех формул пользователя по его айди
func DBFormulaGetByUserID(db *sqlx.DB, userID int64) (*[]entities.Formula, error) {
	formulas := []entities.Formula{}
	query := `SELECT id, title, value, user_id FROM formula WHERE user_id = $1`

	err := db.Select(&formulas, query, userID)
	if err != nil {
		return nil, err
	}
	return &formulas, nil
}

// DBFormulaCreate создание формулы
func DBFormulaCreate(db *sqlx.DB, formula *entities.Formula) (*entities.Formula, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO formula (title, value, user_id)
	VALUES (:title, :value, :user_id) RETURNING id`

	stmt, err := tx.PrepareNamed(query)
	if stmt == nil {
		tx.Rollback()
		return nil, errors.New("error preparing statement")
	}

	err = stmt.Get(&formula.ID, *formula)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	diff, err := util.CompareStrings(formula.Value, formula.Value)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	formulaHistory := entities.FormulaHistory{
		FormulaID:  formula.ID,
		Difference: diff,
		Hash:       util.GenerateHash(formula.Value),
		CodeName:   util.GenerateName(),
	}

	query = `INSERT INTO formula_vcs (formula_id, difference, hash, code_name) VALUES (:formula_id, :difference, :hash, :code_name) RETURNING id`

	stmt, err = tx.PrepareNamed(query)
	if err != nil {
		tx.Rollback()
	}
	err = stmt.Get(&formulaHistory.ID, formulaHistory)

	tx.Commit()
	return formula, nil
}

// DBFormulaUpdate Обновление формулы
func DBFormulaUpdate(db *sqlx.DB, formula *entities.UpdateFormulaRequest) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	oldFormula, err := DBFormulaGetByID(db, int64(formula.ID))
	if err != nil {
		tx.Rollback()
		return err
	}

	query := `UPDATE formula SET title = $1, value = $2 WHERE id = $3`

	_, err = tx.Exec(query, formula.Title, formula.Value, formula.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	diff, err := util.CompareStrings(oldFormula.Value, formula.Value)
	if err != nil {
		tx.Rollback()
		return err
	}
	formulaHistory := entities.FormulaHistory{
		FormulaID:  formula.ID,
		Difference: diff,
		Hash:       util.GenerateHash(formula.Value),
		CodeName:   util.GenerateName(),
	}

	query = `INSERT INTO formula_vcs (formula_id, difference, hash, code_name) VALUES (:formula_id, :difference, :hash, :code_name) RETURNING id`

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		tx.Rollback()
	}
	err = stmt.Get(&formulaHistory.ID, formulaHistory)

	tx.Commit()
	return nil
}

// DBFormulaDelete удаление формулы
func DBFormulaDelete(db *sqlx.DB, formulaID int64) error {
	query := `DELETE FROM formula WHERE id = $1`

	_, err := db.Exec(query, formulaID)
	if err != nil {
		return err
	}
	return nil
}

// DBFormulaHistoryGet получение истории ввода формул
func DBFormulaHistoryGet(db *sqlx.DB, userID int64, pageNumber int64) (*[]entities.Formula, error) {
	formulas := []entities.Formula{}
	query := `
		SELECT id, title, value, user_id FROM formula WHERE user_id = $1
		ORDER BY id DESC
		LIMIT 20
		OFFSET $2
	`

	err := db.Select(&formulas, query, userID, (pageNumber-1)*20)
	if err != nil {
		return nil, err
	}
	return &formulas, nil
}
