package repository

import (
	"fmt"

	platform "github.com/dvd-denis/IT-Platform"
	"github.com/jmoiron/sqlx"
)

type AuthorPostgres struct {
	db *sqlx.DB
}

func NewAuthorPostgres(db *sqlx.DB) *AuthorPostgres {
	return &AuthorPostgres{db: db}
}

func (r *AuthorPostgres) Create(author platform.Authors) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", authorTable)
	row := r.db.QueryRow(query, author.Name)
	err := row.Scan(&id)

	return id, err
}

func (r *AuthorPostgres) GetAllLimit(limit int) ([]platform.Authors, error) {
	var authors []platform.Authors

	query := fmt.Sprintf("SELECT id, name FROM %s LIMIT $1", authorTable)
	err := r.db.Select(&authors, query, limit)

	return authors, err
}

func (r *AuthorPostgres) GetById(id int) (platform.Authors, error) {
	var author platform.Authors

	query := fmt.Sprintf("SELECT id, name FROM %s WHERE id = $1", authorTable)
	err := r.db.Get(&author, query, id)

	return author, err
}

func (r *AuthorPostgres) GetByName(name string) ([]platform.Authors, error) {
	var authors []platform.Authors

	query := fmt.Sprintf("SELECT id, name FROM %s WHERE name ILIKE $1", authorTable)
	err := r.db.Select(&authors, query, name+"%")
	if err != nil {
		return authors, err
	}

	return authors, err
}

func (r *AuthorPostgres) GetPlatformsById(id int) ([]platform.Platforms, error) {
	var platforms []platform.Platforms
	var id_platforms []int

	query := fmt.Sprintf("SELECT platform_id FROM %s WHERE author_id = $1", authorsListsTable)
	if err := r.db.Select(&id_platforms, query, id); err != nil {
		return platforms, err
	}

	var author_name string
	query = fmt.Sprintf("SELECT name FROM %s WHERE id = $1", authorTable)
	if err := r.db.Get(&author_name, query, id); err != nil {
		return platforms, err
	}

	var platform platform.Platforms
	for _, i := range id_platforms {
		query := fmt.Sprintf("SELECT id, title, text FROM %s WHERE id = $1", platformsTable)
		if err := r.db.Get(&platform, query, i); err != nil {
			return platforms, err
		}
		platform.Author = author_name
		platforms = append(platforms, platform)
	}

	return platforms, nil
}

func (r *AuthorPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s al USING %s pm WHERE al.author_id = $1 AND al.platform_id = pm.id", authorsListsTable, platformsTable)
	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE id=$1", authorTable)
	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (r *AuthorPostgres) Update(id int, input platform.UpdateAuthorInput) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id = $2", authorTable)
	_, err := r.db.Exec(query, input.Name, id)
	return err
}
