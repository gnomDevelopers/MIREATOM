package postgres

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"server/internal/entities"
)

func DBArticleCreate(db *sqlx.DB, article *entities.Article) (*entities.Article, error) {
	query := `INSERT INTO articles (user_id, title, science, section, path)
	VALUES (:user_id, :title, :science, :section, :path) RETURNING id`

	stmt, err := db.PrepareNamed(query)
	if stmt == nil {
		return nil, errors.New("error preparing statement")
	}
	err = stmt.Get(&article.Id, article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func DBArticleGetAll(db *sqlx.DB) (*[]entities.Article, error) {
	articles := []entities.Article{}
	query := `SELECT id, user_id, title, id, user_id, title, science, section, path FROM articles`

	err := db.Select(&articles, query)
	if err != nil {
		return nil, err
	}
	return &articles, nil
}

func DBArticleGetPath(db *sqlx.DB, id int) (string, error) {
	var path string
	query := `SELECT path FROM articles WHERE id = $1`
	err := db.Get(&path, query, id)
	if err != nil {
		return "", err
	}

	return path, nil
}

func DBArticleUpdate(db *sqlx.DB, article *entities.UpdateArticleRequest) error {
	query := `UPDATE articles SET title = $1, science = $2, section = $3 WHERE id = $4`

	_, err := db.Exec(query, article.Title, article.Science, article.Section, article.Id)
	if err != nil {
		return err
	}
	return nil
}

func DBArticleUpdatePath(db *sqlx.DB, path string, id int) error {
	query := `UPDATE articles SET path = $1 WHERE id = $2`

	_, err := db.Exec(query, path, id)
	if err != nil {
		return err
	}
	return nil
}

func DBArticleDelete(db *sqlx.DB, articleId int64) error {
	query := `DELETE FROM articles WHERE id = $1`

	_, err := db.Exec(query, articleId)
	if err != nil {
		return err
	}
	return nil
}
