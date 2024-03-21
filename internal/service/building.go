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
