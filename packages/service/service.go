package service

import (
	"github.com/dvd-denis/poem-app"
	"github.com/dvd-denis/poem-app/packages/repository"
)

type Authorization interface {
	CreateUser(user poem.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Poem interface {
	Create(authorId int, poem poem.Poems) (int, error)
}

type Service struct {
	Authorization
	Poem
}

func NewService(repos repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Poem:          newPoemService(repos.Poem),
	}
}
