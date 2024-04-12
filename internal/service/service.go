package service

import (
	"github.com/TandDA/coursedb/internal/model"
	"github.com/TandDA/coursedb/internal/repository"
)

type Service struct {
	Building
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Building: NewBuildingService(r.Building),
	}
}

type Building interface {
	Save(b model.Building) (string, error)
	GetAll() ([]model.Building, error)
	Update(bld model.Building) error
	Delete(id string) error
}
