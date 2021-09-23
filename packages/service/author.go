package service

import (
	platform "github.com/dvd-denis/IT-Platform"
	"github.com/dvd-denis/IT-Platform/packages/repository"
)

type AuthorService struct {
	repo repository.Author
}

func newAuthorService(repo repository.Author) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) Create(author platform.Authors) (int, error) {
	return s.repo.Create(author)
}

func (s *AuthorService) GetById(id int) (platform.Authors, error) {
	return s.repo.GetById(id)
}

func (s *AuthorService) GetByName(name string) ([]platform.Authors, error) {
	return s.repo.GetByName(name)
}

func (s *AuthorService) GetAllLimit(limit int) ([]platform.Authors, error) {
	return s.repo.GetAllLimit(limit)
}

func (s *AuthorService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *AuthorService) Update(id int, input platform.UpdateAuthorInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}

func (s *AuthorService) GetPlatformsById(id int) ([]platform.Platforms, error) {
	return s.repo.GetPlatformsById(id)
}
