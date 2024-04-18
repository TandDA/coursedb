package service

import (
	"github.com/TandDA/coursedb/internal/model"
	"github.com/TandDA/coursedb/internal/repository"
)

type GuestService struct {
	repo repository.Guest
}

func NewGuestService(repo repository.Guest) *GuestService {
	return &GuestService{repo: repo}
}

func (s *GuestService) GetAllGuestsWithComplains() ([]model.GuestAndComplain, error) {
	return s.repo.GetAllGuestsWithComplains()
}

func (s *GuestService) GetAll(from, to string) ([]model.Guest, error) {
	return s.repo.GetAll(from, to)
}
