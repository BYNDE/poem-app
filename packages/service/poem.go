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
