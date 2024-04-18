package service

import (
	"github.com/TandDA/coursedb/internal/model"
	"github.com/TandDA/coursedb/internal/repository"
)

type Service struct {
	Building
	Room
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Building: NewBuildingService(r.Building),
		Room:     NewRoomService(r.Room),
	}
}

type Building interface {
	Save(b model.Building) (string, error)
	GetAll() ([]model.Building, error)
	Update(bld model.Building) error
	Delete(id string) error
}

type Room interface {
	GetAllFreeRooms(floorNumber, class, numberOfRooms int) ([]model.Room, error)
}
