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

type Repository struct {
	Building
	Room
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Building: NewBuildingRepository(db),
		Room:     NewRoomRepository(db),
	}
}
