package repository

import (
	"fmt"
	"strings"

	platform "github.com/dvd-denis/IT-Platform"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PlatformPostgres struct {
	db *sqlx.DB
}

func NewPlatformPostgres(db *sqlx.DB) *PlatformPostgres {
	return &PlatformPostgres{db: db}
}

func (r *PlatformPostgres) Create(authorId int, platform platform.Platforms) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createPlatformQuery := fmt.Sprintf("INSERT INTO %s (title, text) VALUES ($1, $2) RETURNING id", platformsTable)
	row := tx.QueryRow(createPlatformQuery, platform.Title, platform.Text)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	CreateAuthorsListQuery := fmt.Sprintf("INSERT INTO %s (author_id, platform_id) VALUES ($1, $2)", authorsListsTable)
	_, err = tx.Exec(CreateAuthorsListQuery, authorId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()

}

func (r *PlatformPostgres) GetAllLimit(limit int) ([]platform.Platforms, error) {
	var platforms []platform.Platforms

	query := fmt.Sprintf("SELECT id, title, text FROM %s LIMIT $1", platformsTable)
	if err := r.db.Select(&platforms, query, limit); err != nil {
		return platforms, err
	}

	var authorId int

	for i := range platforms {
		query = fmt.Sprintf("SELECT author_id FROM %s WHERE platform_id = $1", authorsListsTable)
		if err := r.db.Get(&authorId, query, platforms[i].Id); err != nil {
			return platforms, err
		}

		query = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", authorTable)
		if err := r.db.Get(&platforms[i].Author, query, authorId); err != nil {
			return platforms, err
		}
	}

	return platforms, nil
}

func (r *PlatformPostgres) GetById(id int) (platform.Platforms, error) {
	var platform platform.Platforms

	query := fmt.Sprintf("SELECT id, title, text FROM %s WHERE id = $1", platformsTable)
	if err := r.db.Get(&platform, query, id); err != nil {
		return platform, err
	}

	var authorId int

	query = fmt.Sprintf("SELECT author_id FROM %s WHERE platform_id = $1", authorsListsTable)
	if err := r.db.Get(&authorId, query, id); err != nil {
		return platform, err
	}

	query = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", authorTable)
	err := r.db.Get(&platform.Author, query, authorId)

	return platform, err
}

func (r *PlatformPostgres) GetByTitle(title string) ([]platform.Platforms, error) {
	var platforms []platform.Platforms

	query := fmt.Sprintf("SELECT id, title, text FROM %s WHERE title ILIKE $1", platformsTable)
	if err := r.db.Select(&platforms, query, title+"%"); err != nil {
		return platforms, err
	}

	var authorId int

	for i := range platforms {
		query = fmt.Sprintf("SELECT author_id FROM %s WHERE platform_id = $1", authorsListsTable)
		if err := r.db.Get(&authorId, query, platforms[i].Id); err != nil {
			return platforms, err
		}

		query = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", authorTable)
		if err := r.db.Get(&platforms[i].Author, query, authorId); err != nil {
			return platforms, err
		}
	}

	return platforms, nil
}

func (r *PlatformPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s pm USING %s al WHERE pm.id = al.platform_id AND al.platform_id=$1", platformsTable, authorsListsTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *PlatformPostgres) Update(id int, input platform.UpdatePlatformInput) error {
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

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", platformsTable, setQuery, argId)

	args = append(args, id)

	logrus.Debugf("updatedQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err

}
