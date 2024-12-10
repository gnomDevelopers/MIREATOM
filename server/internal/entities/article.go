package entities

type Article struct {
	Id      int    `json:"id" db:"id"`
	UserId  int    `json:"user_id" db:"user_id"`
	Title   string `json:"title" db:"title"`
	Science string `json:"science" db:"science"`
	Section string `json:"section" db:"section"`
	Path    string `json:"path" db:"path"`
}

type CreateArticleRequest struct {
	Title   string `form:"title"`
	Science string `form:"science"`
	Section string `form:"section"`
	File    string `form:"file"`
}
