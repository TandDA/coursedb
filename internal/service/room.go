package service

import (
	"github.com/TandDA/coursedb/internal/model"
	"github.com/TandDA/coursedb/internal/repository"
)

type RoomService struct {
	repo repository.Room
}

func NewRoomService(repo repository.Room) *RoomService {
	return &RoomService{repo: repo}
}

func (r RoomService) GetAllFreeRooms(floorNumber, class, numberOfRooms int) ([]model.Room, error) {
	return r.repo.GetAllFreeRooms(floorNumber, class, numberOfRooms)
}
