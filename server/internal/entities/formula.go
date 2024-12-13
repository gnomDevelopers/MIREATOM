package entities

type Formula struct {
	ID     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Value  string `json:"value" db:"value"`
	UserID int    `json:"user_id" db:"user_id"`
}

type CreateFormulaRequest struct {
	Value string `json:"value" db:"value"`
	Title string `json:"title" db:"title"`
}

type CreateFormulaResponse struct {
	ID int `json:"id" db:"id"`
}

type UpdateFormulaRequest struct {
	ID    int    `json:"id" db:"id"`
	Value string `json:"value" db:"value"`
	Title string `json:"title" db:"title"`
}

type GetFormulaFromArticleResponse struct {
	Formula string `json:"formula"`
}
