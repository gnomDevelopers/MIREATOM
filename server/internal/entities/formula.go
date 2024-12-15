package entities

// Formula базовая структура формулы
type Formula struct {
	ID     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Value  string `json:"value" db:"value"`
	UserID int    `json:"user_id" db:"user_id"`
}

// CreateFormulaRequest структура формулы для запроса создания
type CreateFormulaRequest struct {
	Value string `json:"value" db:"value"`
	Title string `json:"title" db:"title"`
}

// CreateFormulaResponse структура формулы для ответа создания
type CreateFormulaResponse struct {
	ID int `json:"id" db:"id"`
}

// UpdateFormulaRequest структура формулы для запроса обновления
type UpdateFormulaRequest struct {
	ID    int    `json:"id" db:"id"`
	Value string `json:"value" db:"value"`
	Title string `json:"title" db:"title"`
}

// GetFormulaFromArticleResponse структура формулы для получения формулы из статьи
type GetFormulaFromArticleResponse struct {
	Formula string `json:"formula"`
}

// RecognizedFormula структура формулы, распознанной из картинки
type RecognizedFormula struct {
	Formula string `json:"formula"`
}
