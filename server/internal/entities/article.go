package entities

type Article struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Title   string `json:"title"`
	Science string `json:"science"`
	Section string `json:"section"`
	Path    string `json:"path"`
}
