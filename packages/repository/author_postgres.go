package repository

import (
	"fmt"

	"github.com/dvd-denis/poem-app"
	"github.com/jmoiron/sqlx"
)

type AuthorPostgres struct {
	db *sqlx.DB
}

func NewAuthorPostgres(db *sqlx.DB) *AuthorPostgres {
	return &AuthorPostgres{db: db}
}

func (r *AuthorPostgres) Create(author poem.Authors) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", authorTable)
	row := r.db.QueryRow(query, author.Name)
	err := row.Scan(&id)

	return id, err
}

func (r *AuthorPostgres) GetAllLimit(limit int) ([]poem.Authors, error) {
	var authors []poem.Authors

	query := fmt.Sprintf("SELECT id, name FROM %s LIMIT $1", authorTable)
	err := r.db.Select(&authors, query, limit)

	return authors, err
}

func (r *AuthorPostgres) GetById(id int) (poem.Authors, error) {
	var author poem.Authors

	query := fmt.Sprintf("SELECT id, name FROM %s WHERE id = $1", authorTable)
	err := r.db.Get(&author, query, id)

	return author, err
}

func (r *AuthorPostgres) GetByName(name string) ([]poem.Authors, error) {
	var authors []poem.Authors

	query := fmt.Sprintf("SELECT id, name FROM %s WHERE name LIKE $1", authorTable)
	err := r.db.Select(&authors, query, name+"%")
	if err != nil {
		return authors, err
	}

	return authors, err
}

func (r *AuthorPostgres) GetPoemsById(id int) ([]poem.Poems, error) {
	var poems []poem.Poems
	var id_poems []int

	query := fmt.Sprintf("SELECT poem_id FROM %s WHERE author_id = $1", authorsListsTable)
	if err := r.db.Select(&id_poems, query, id); err != nil {
		return poems, err
	}

	var author_name string
	query = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", authorTable)
	if err := r.db.Get(&author_name, query, id); err != nil {
		return poems, err
	}

	var poem poem.Poems
	for _, i := range id_poems {
		query := fmt.Sprintf("SELECT id, title, text FROM %s WHERE id = $1", poemsTable)
		if err := r.db.Get(&poem, query, i); err != nil {
			return poems, err
		}
		poem.Author = author_name
		poems = append(poems, poem)
	}

	return poems, nil
}

func (r *AuthorPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s al USING %s pm WHERE al.author_id = $1 AND al.poem_id = pm.id", authorsListsTable, poemsTable)
	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE id=$1", authorTable)
	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (r *AuthorPostgres) Update(id int, input poem.UpdateAuthorInput) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id = $2", authorTable)
	_, err := r.db.Exec(query, input.Name, id)
	return err
}
