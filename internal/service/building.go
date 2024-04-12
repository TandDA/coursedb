package service

import (
	"github.com/TandDA/coursedb/internal/model"
	"github.com/TandDA/coursedb/internal/repository"
)

type BuildingService struct {
	repo repository.Building
}

func NewBuildingService(repo repository.Building) *BuildingService {
	return &BuildingService{repo: repo}
}

func (s *BuildingService) Save(b model.Building) (string, error) {
	return s.repo.Save(b)
}

func (s *BuildingService) GetAll() ([]model.Building, error) {
	return s.repo.GetAll()
}

func (s *BuildingService) Update(bld model.Building) error {
	return s.repo.Update(bld)
}

func (s *BuildingService) Delete(id string) error {
	return s.repo.Delete(id)
}