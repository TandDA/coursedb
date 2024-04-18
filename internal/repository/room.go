package repository

import (
	"database/sql"

	"github.com/TandDA/coursedb/internal/model"
)

type RoomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) GetAllFreeRooms() ([]model.Room, error) {
	return nil, nil
}
