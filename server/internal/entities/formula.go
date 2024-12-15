package entities

type Formula struct {
	ID     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Value  string `json:"value" db:"value"`
	UserID int    `json:"user_id" db:"user_id"`
}

type FormulaInfo struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Value    string `json:"value" db:"value"`
	UserID   int    `json:"user_id" db:"user_id"`
	FullName string `json:"full_name" db:"full_name"`
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

type RecognizedFormula struct {
	Formula string `json:"formula"`
}

type FormulaAnalysisRequest struct {
	Formula string `json:"formula"`
}

type FormulaAnalysisResponse struct {
	Percent      string `json:"percent"`
	MatchFormula string `json:"match_formula"`
	Author       string `json:"author"`
	Name         string `json:"name"`
}
