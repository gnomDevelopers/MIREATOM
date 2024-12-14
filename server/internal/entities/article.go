package entities

type Article struct {
	Id      int    `json:"id" db:"id"`
	UserId  int    `json:"user_id" db:"user_id"`
	Title   string `json:"title" db:"title"`
	Science string `json:"science" db:"science"`
	Section string `json:"section" db:"section"`
	Path    string `json:"path" db:"path"`
}

type ArticleInfo struct {
	Id       int    `json:"id" db:"article_id"`
	UserId   int    `json:"user_id" db:"user_id"`
	Title    string `json:"title" db:"title"`
	Science  string `json:"science" db:"science"`
	Section  string `json:"section" db:"section"`
	FullName string `json:"full_name" db:"full_name"`
}

type CreateArticleRequest struct {
	Title   string `form:"title"`
	Science string `form:"science"`
	Section string `form:"section"`
	File    string `form:"file"`
}

type UpdateArticleRequest struct {
	Id      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Science string `json:"science" db:"science"`
	Section string `json:"section" db:"section"`
}
