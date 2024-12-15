package postgres

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"server/internal/entities"
)

// DBArticleExists проверка на существование статьи в бд
func DBArticleExists(db *sqlx.DB, title string, userId int) (bool, error) {
	exists := 0
	query := `SELECT 1 FROM articles WHERE title = $1 AND user_id = $2 LIMIT 1`

	err := db.QueryRow(query, title, userId).Scan(&exists)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

// DBArticleCreate создание статьи
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

// DBArticleGetAll получение всех статей
func DBArticleGetAll(db *sqlx.DB) (*[]entities.ArticleInfo, error) {
	articles := []entities.ArticleInfo{}
	query := `
		SELECT articles.id AS article_id, articles.title, articles.science, articles.section, users.id AS user_id,
    	CONCAT(users.surname, ' ', users.name, ' ', users.third_name) AS full_name
		FROM articles
		JOIN users ON articles.user_id = users.id;
	`

	err := db.Select(&articles, query)
	if err != nil {
		return nil, err
	}
	return &articles, nil
}

// DBArticleGetByUserId получение всех статей одного пользователя
func DBArticleGetByUserId(db *sqlx.DB, id int) (*[]entities.ArticleInfo, error) {
	articles := []entities.ArticleInfo{}
	query := `
		SELECT articles.id AS article_id, articles.title, articles.science, articles.section, users.id AS user_id,
    	CONCAT(users.surname, ' ', users.name, ' ', users.third_name) AS full_name
		FROM articles
		JOIN users ON articles.user_id = users.id WHERE user_id = $1;
	`

	err := db.Select(&articles, query, id)
	if err != nil {
		return nil, err
	}
	return &articles, nil
}

// DBArticleGetPath получение пути сохраненной статьи
func DBArticleGetPath(db *sqlx.DB, id int) (string, error) {
	var path string
	query := `SELECT path FROM articles WHERE id = $1`
	err := db.Get(&path, query, id)
	if err != nil {
		return "", err
	}

	return path, nil
}

// DBArticleUpdate обновление статьи
func DBArticleUpdate(db *sqlx.DB, article *entities.UpdateArticleRequest) error {
	query := `UPDATE articles SET title = $1, science = $2, section = $3 WHERE id = $4`

	_, err := db.Exec(query, article.Title, article.Science, article.Section, article.Id)
	if err != nil {
		return err
	}
	return nil
}

// DBArticleUpdatePath обновление пути сохраненной статьи
func DBArticleUpdatePath(db *sqlx.DB, path string, id int) error {
	query := `UPDATE articles SET path = $1 WHERE id = $2`

	_, err := db.Exec(query, path, id)
	if err != nil {
		return err
	}
	return nil
}

// DBArticleDelete  удаление статьи
func DBArticleDelete(db *sqlx.DB, articleId int64) error {
	query := `DELETE FROM articles WHERE id = $1`

	_, err := db.Exec(query, articleId)
	if err != nil {
		return err
	}
	return nil
}
