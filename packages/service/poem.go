package service

import (
	"github.com/dvd-denis/poem-app"
	"github.com/dvd-denis/poem-app/packages/repository"
)

type PoemService struct {
	repo repository.Poem
}

func newPoemService(repo repository.Poem) *PoemService {
	return &PoemService{repo: repo}
}

func (s *PoemService) Create(authorId int, poem poem.Poems) (int, error) {
	return s.repo.Create(authorId, poem)
}

func (s *PoemService) GetAllLimit(limit int) ([]poem.Poems, error) {
	return s.repo.GetAllLimit(limit)
}

func (s *PoemService) GetById(id int) (poem.Poems, error) {
	return s.repo.GetById(id)
}

func (s *PoemService) GetByTitle(title string) ([]poem.Poems, error) {
	return s.repo.GetByTitle(title)
}

func (s *PoemService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *PoemService) Update(id int, input poem.UpdatePoemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}
