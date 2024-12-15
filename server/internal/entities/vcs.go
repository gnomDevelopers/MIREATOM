package entities

// FormulaHistory структура истории изменения формулы
type FormulaHistory struct {
	ID         int    `json:"id" db:"id"`
	FormulaID  int    `json:"formula_id" db:"formula_id"`
	Difference string `json:"difference" db:"difference"`
	Hash       string `json:"hash" db:"hash"`
	CodeName   string `json:"code_name" db:"code_name"`
	CreatedAt  string `json:"created_at" db:"created_at"`
}
