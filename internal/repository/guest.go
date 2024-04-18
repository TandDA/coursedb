package repository

import (
	"database/sql"
	"github.com/TandDA/coursedb/internal/model"
)

type GuestRepository struct {
	db *sql.DB
}

func NewGuestRepository(db *sql.DB) *GuestRepository {
	return &GuestRepository{db: db}
}

func (r *GuestRepository) GetAllGuestsWithComplains() ([]model.GuestAndComplain, error) {
	query := `SELECT g.*, c.id, c.complain_text FROM guest g
JOIN complain c ON c.guest_id = g.id `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ans []model.GuestAndComplain
	for rows.Next() {
		var rm model.GuestAndComplain
		err := rows.Scan(&rm.Id, &rm.FirstName, &rm.LastName, &rm.DateOfEntry, &rm.Complain.Id, &rm.ComplainText)
		if err != nil {
			return nil, err
		}
		ans = append(ans, rm)
	}
	return ans, nil
}
