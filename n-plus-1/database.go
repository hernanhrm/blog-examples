package n_plus_1

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DBConnection() (*sql.DB, error) {
	dns := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"postgres",
		"secret",
		"localhost",
		5432,
		"restaurant",
	)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
