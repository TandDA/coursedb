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
	query := `
		SELECT r.*
		FROM room r
		LEFT JOIN booking b ON r.id = b.room_id AND CURRENT_DATE BETWEEN b.date_of_entry AND b.date_of_departure
		WHERE b.id IS NULL;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ans []model.Room
	for rows.Next() {
		var rm model.Room
		err := rows.Scan(&rm.Id, &rm.NumberOfRooms, &rm.RegularPrice, &rm.FloorId)
		if err != nil {
			return nil, err
		}
		ans = append(ans, rm)
	}
	return ans, nil
}
