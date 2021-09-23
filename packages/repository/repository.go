package repository

import (
	platform "github.com/dvd-denis/IT-Platform"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user platform.User) (int, error)
	GetUser(username, password string) (platform.User, error)
}

type Platform interface {
	Create(authorId int, platform platform.Platforms) (int, error)
	GetById(id int) (platform.Platforms, error)
	GetByTitle(title string) ([]platform.Platforms, error)
	Delete(id int) error
	Update(id int, input platform.UpdatePlatformInput) error
	GetAllLimit(limit int) ([]platform.Platforms, error)
}

type Author interface {
	Create(author platform.Authors) (int, error)
	GetById(id int) (platform.Authors, error)
	GetByName(name string) ([]platform.Authors, error)
	Delete(id int) error
	Update(id int, input platform.UpdateAuthorInput) error
	GetPlatformsById(id int) ([]platform.Platforms, error)
	GetAllLimit(limit int) ([]platform.Authors, error)
}

type Repository struct {
	Authorization
	Platform
	Author
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Platform:      NewPlatformPostgres(db),
		Author:        NewAuthorPostgres(db),
	}
}
