package postgres

import (
	"github.com/jmoiron/sqlx"
	"server/internal/entities"
)

// DBGetFormulaCommits получение истории изменений определенной формулы по ее айди
func DBGetFormulaCommits(db *sqlx.DB, formulaID int64) (*[]entities.FormulaHistory, error) {
	formulaHistory := []entities.FormulaHistory{}
	query := `
		SELECT * FROM formula_vcs WHERE formula_id = $1
		ORDER BY created_at DESC
	`

	err := db.Select(&formulaHistory, query, formulaID)
	if err != nil {
		return nil, err
	}
	return &formulaHistory, nil
}
