package service

import (
	"github.com/dvd-denis/poem-app"
	"github.com/dvd-denis/poem-app/packages/repository"
)

type AuthorService struct {
	repo repository.Author
}

func newAuthorService(repo repository.Author) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) Create(author poem.Authors) (int, error) {
	return s.repo.Create(author)
}

func (s *AuthorService) GetById(id int) (poem.Authors, error) {
	return s.repo.GetById(id)
}

func (s *AuthorService) GetByName(name string) ([]poem.Authors, error) {
	return s.repo.GetByName(name)
}

func (s *AuthorService) GetAllLimit(limit int) ([]poem.Authors, error) {
	return s.repo.GetAllLimit(limit)
}

func (s *AuthorService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *AuthorService) Update(id int, input poem.UpdateAuthorInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}

func (s *AuthorService) GetPoemsById(id int) ([]poem.Poems, error) {
	return s.repo.GetPoemsById(id)
}
