package service

import (
	"github.com/TandDA/coursedb/internal/model"
	"github.com/TandDA/coursedb/internal/repository"
)

type Service struct {
	Building
}

func NewService(r *repository.Repository) *Service {
	return &Service{}
}

type Building interface {
	Save(b model.Building) (string, error)
}