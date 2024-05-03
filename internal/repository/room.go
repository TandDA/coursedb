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

func (r *RoomRepository) GetAllFreeRooms(floorNumber, class, numberOfRooms int) ([]model.Room, error) {
	query := `
		SELECT r.*
FROM room r
LEFT JOIN booking b ON r.id = b.room_id AND CURRENT_DATE BETWEEN b.date_of_entry AND b.date_of_departure
JOIN floor f ON r.floor_id = f.id 
JOIN building bld ON f.building_id = bld.id 
WHERE b.id IS NULL 
AND ($1 = -1 OR floor_number = $1)
AND ($2 = -1 OR class = $2)
AND ($3 = -1 OR number_of_rooms = $3);`

	rows, err := r.db.Query(query, floorNumber, class, numberOfRooms)
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
