package repository

import (
	"database/sql"
	"github.com/TandDA/coursedb/internal/model"
)

type Building interface {
	Save(b model.Building) (string, error)
	GetAll() ([]model.Building, error)
	Update(bld model.Building) error
	Delete(id string) error
}

type Room interface {
	GetAllFreeRooms(floorNumber, class, numberOfRooms int) ([]model.Room, error)
}

type Guest interface {
	GetAll(from, to string) ([]model.Guest, error)
	GetAllGuestsWithComplains() ([]model.GuestAndComplain, error)
}

type Repository struct {
	Building
	Room
	Guest
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Building: NewBuildingRepository(db),
		Room:     NewRoomRepository(db),
		Guest:    NewGuestRepository(db),
	}
}
