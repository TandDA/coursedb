package repository

import (
	"database/sql"
	"github.com/TandDA/coursedb/internal/model"
	"time"
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

func (r *GuestRepository) GetAll(from, to string) ([]model.Guest, error) {
	// Парсим строки в формат даты
	f, err := time.Parse("2006.01.02", from)
	if err != nil {
		return nil, err
	}
	t, err := time.Parse("2006.01.02", to)
	if err != nil {
		return nil, err
	}

	// Форматируем даты для SQL запроса
	fromDate := f.Format("2006-01-02")
	toDate := t.Format("2006-01-02")
	query := `SELECT g.* FROM guest g
WHERE ($1::date IS NULL OR $2::date IS NULL OR date_of_entry BETWEEN $1::date AND $2::date)`

	rows, err := r.db.Query(query, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ans []model.Guest
	for rows.Next() {
		var rm model.Guest
		err := rows.Scan(&rm.Id, &rm.FirstName, &rm.LastName, &rm.DateOfEntry)
		if err != nil {
			return nil, err
		}
		ans = append(ans, rm)
	}
	return ans, nil
}
