package repository

import (
	"database/sql"

	"github.com/TandDA/coursedb/internal/model"
	"github.com/google/uuid"
)

type BuildingRepository struct {
	db *sql.DB
}

func NewBuildingRepository(db *sql.DB) *BuildingRepository {
	return &BuildingRepository{db: db}
}

func (r *BuildingRepository) Save(b model.Building) (string, error) {
	uuid := uuid.New()
	query := "INSERT INTO building(id, class, number_of_floors, address) VALUES(?,?,?,?)"
	_, err := r.db.Exec(query, uuid.String(), b.Class, b.NumberOfFloors, b.Address)
	return uuid.String(), err
}