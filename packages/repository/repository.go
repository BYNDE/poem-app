package repository

import (
	"github.com/dvd-denis/poem-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user poem.User) (int, error)
	GetUser(username, password string) (poem.User, error)
}

type Poem interface {
	Create(authorId int, poem poem.Poems) (int, error)
}

type Repository struct {
	Authorization
	Poem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Poem:          NewPoemPostgres(db),
	}
}
