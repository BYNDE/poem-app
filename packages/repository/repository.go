package repository

import (
	"github.com/dvd-denis/IT-Platform"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user poem.User) (int, error)
	GetUser(username, password string) (poem.User, error)
}

type Poem interface {
	Create(authorId int, poem poem.Poems) (int, error)
	GetById(id int) (poem.Poems, error)
	GetByTitle(title string) ([]poem.Poems, error)
	Delete(id int) error
	Update(id int, input poem.UpdatePoemInput) error
	GetAllLimit(limit int) ([]poem.Poems, error)
}

type Author interface {
	Create(author poem.Authors) (int, error)
	GetById(id int) (poem.Authors, error)
	GetByName(name string) ([]poem.Authors, error)
	Delete(id int) error
	Update(id int, input poem.UpdateAuthorInput) error
	GetPoemsById(id int) ([]poem.Poems, error)
	GetAllLimit(limit int) ([]poem.Authors, error)
}

type Repository struct {
	Authorization
	Poem
	Author
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Poem:          NewPoemPostgres(db),
		Author:        NewAuthorPostgres(db),
	}
}
