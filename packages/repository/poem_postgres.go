package repository

import (
	"fmt"

	"github.com/dvd-denis/poem-app"
	"github.com/jmoiron/sqlx"
)

type PoemPostgres struct {
	db *sqlx.DB
}

func NewPoemPostgres(db *sqlx.DB) *PoemPostgres {
	return &PoemPostgres{db: db}
}

func (r *PoemPostgres) Create(poem poem.Poems) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createPoemQuery := fmt.Sprintf("INSERT INTO %s (title, text, author_id) VALUES ($1, $2, $3) RETURNING id", poemsTable)
	row := tx.QueryRow(createPoemQuery, poem.Title, poem.Text, poem.AuthorId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()

}
