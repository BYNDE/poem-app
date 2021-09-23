package service

import (
	platform "github.com/dvd-denis/IT-Platform"
	"github.com/dvd-denis/IT-Platform/packages/repository"
)

type Authorization interface {
	CreateUser(user platform.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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
	GetAllLimit(limit int) ([]platform.Authors, error)
	GetById(id int) (platform.Authors, error)
	GetByName(name string) ([]platform.Authors, error)
	Delete(id int) error
	Update(id int, input platform.UpdateAuthorInput) error
	GetPlatformsById(id int) ([]platform.Platforms, error)
}

type Service struct {
	Authorization
	Platform
	Author
}

func NewService(repos repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Platform:      newPlatformService(repos.Platform),
		Author:        newAuthorService(repos.Author),
	}
}
