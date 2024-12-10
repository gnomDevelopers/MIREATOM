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
