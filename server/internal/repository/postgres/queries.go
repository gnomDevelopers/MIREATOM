package postgres

const (
	createUserTable = `
		CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL,
		surname VARCHAR NOT NULL,
		third_name VARCHAR,
		role VARCHAR NOT NULL,
		email VARCHAR NOT NULL UNIQUE,
		password VARCHAR NOT NULL
);
`

	createArticleTable = `
		CREATE TABLE IF NOT EXISTS articles (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		title VARCHAR NOT NULL,
		path VARCHAR NOT NULL,
		originality VARCHAR[] NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users (id)
);
`

	createArticleVCSTable = `
		CREATE TABLE IF NOT EXISTS article_vcs (
		id SERIAL PRIMARY KEY,
		article_id INT NOT NULL,
		difference VARCHAR NOT NULL,
		hash VARCHAR NOT NULL,
		FOREIGN KEY (article_id) REFERENCES articles (id)
);
	`

	createTableFormula = `
		CREATE TABLE IF NOT EXISTS formula (
		id SERIAL PRIMARY KEY,
		title VARCHAR NOT NULL,
		value VARCHAR NOT NULL,
		user_id INT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users (id)
);
	`
)
