package service

import (
	platform "github.com/dvd-denis/IT-Platform"
	"github.com/dvd-denis/IT-Platform/packages/repository"
)

type PlatformService struct {
	repo repository.Platform
}

func newPlatformService(repo repository.Platform) *PlatformService {
	return &PlatformService{repo: repo}
}

func (s *PlatformService) Create(authorId int, platform platform.Platforms) (int, error) {
	return s.repo.Create(authorId, platform)
}

func (s *PlatformService) GetAllLimit(limit int) ([]platform.Platforms, error) {
	return s.repo.GetAllLimit(limit)
}

func (s *PlatformService) GetById(id int) (platform.Platforms, error) {
	return s.repo.GetById(id)
}

func (s *PlatformService) GetByTitle(title string) ([]platform.Platforms, error) {
	return s.repo.GetByTitle(title)
}

func (s *PlatformService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *PlatformService) Update(id int, input platform.UpdatePlatformInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}
