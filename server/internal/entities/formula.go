package entities

type Formula struct {
	ID     int    `json:"id" db:"id"`
	Value  string `json:"value" db:"value"`
	UserID string `json:"user_id" db:"user_id"`
}

type CreateFormulaRequest struct {
	Value  string `json:"value" db:"value"`
	UserID string `json:"user_id" db:"user_id"`
}

type CreateFormulaResponse struct {
	ID int `json:"id" db:"id"`
}

type UpdateFormulaRequest struct {
	Value string `json:"value" db:"value"`
}

type UpdateFormulaResponse struct {
	ID int `json:"id" db:"id"`
}
