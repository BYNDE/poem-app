package service

import (
	"github.com/dvd-denis/IT-Platform"
	"github.com/dvd-denis/IT-Platform/packages/repository"
)

type Authorization interface {
	CreateUser(user poem.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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
	GetAllLimit(limit int) ([]poem.Authors, error)
	GetById(id int) (poem.Authors, error)
	GetByName(name string) ([]poem.Authors, error)
	Delete(id int) error
	Update(id int, input poem.UpdateAuthorInput) error
	GetPoemsById(id int) ([]poem.Poems, error)
}

type Service struct {
	Authorization
	Poem
	Author
}

func NewService(repos repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Poem:          newPoemService(repos.Poem),
		Author:        newAuthorService(repos.Author),
	}
}
