package repository

import (
	"fmt"
	"strings"

	"github.com/dvd-denis/poem-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PoemPostgres struct {
	db *sqlx.DB
}

func NewPoemPostgres(db *sqlx.DB) *PoemPostgres {
	return &PoemPostgres{db: db}
}

func (r *PoemPostgres) Create(authorId int, poem poem.Poems) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createPoemQuery := fmt.Sprintf("INSERT INTO %s (title, text) VALUES ($1, $2) RETURNING id", poemsTable)
	row := tx.QueryRow(createPoemQuery, poem.Title, poem.Text)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	CreateAuthorsListQuery := fmt.Sprintf("INSERT INTO %s (author_id, poem_id) VALUES ($1, $2)", authorsListsTable)
	_, err = tx.Exec(CreateAuthorsListQuery, authorId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()

}

func (r *PoemPostgres) GetById(id int) (poem.Poems, error) {
	var poem poem.Poems

	query := fmt.Sprintf("SELECT id, title, text FROM %s WHERE id = $1", poemsTable)
	if err := r.db.Get(&poem, query, id); err != nil {
		return poem, err
	}

	var authorId int

	query = fmt.Sprintf("SELECT author_id FROM %s WHERE poem_id = $1", authorsListsTable)
	if err := r.db.Get(&authorId, query, id); err != nil {
		return poem, err
	}

	query = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", authorTable)
	err := r.db.Get(&poem.Author, query, authorId)

	return poem, err
}

func (r *PoemPostgres) GetByTitle(title string) ([]poem.Poems, error) {
	var poems []poem.Poems

	query := fmt.Sprintf("SELECT id, title, text FROM %s WHERE title LIKE $1", poemsTable)
	err := r.db.Select(&poems, query, title+"%")
	if err != nil {
		return poems, err
	}

	var authorId int

	for i, _ := range poems {
		query = fmt.Sprintf("SELECT author_id FROM %s WHERE poem_id = $1", authorsListsTable)
		err = r.db.Get(&authorId, query, poems[i].Id)
		if err != nil {
			return poems, err
		}

		query = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", authorTable)
		err = r.db.Get(&poems[i].Author, query, authorId) // Выдаёт в utf-8, когда нужно WINDOWS-1251

		if err != nil {
			return poems, err
		}
	}

	return poems, err
}

func (r *PoemPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s pm USING %s al WHERE pm.id = al.poem_id AND al.poem_id=$1", poemsTable, authorsListsTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *PoemPostgres) Update(id int, input poem.UpdatePoemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argId))
		args = append(args, *input.Text)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", poemsTable, setQuery, argId)

	args = append(args, id)

	logrus.Debugf("updatedQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err

}
