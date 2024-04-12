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
	query := "INSERT INTO building(id, class, number_of_floors, address) VALUES($1,$2,$3,$4)"
	_, err := r.db.Exec(query, uuid.String(), b.Class, b.NumberOfFloors, b.Address)
	return uuid.String(), err
}

func (r *BuildingRepository) GetAll() ([]model.Building, error) {
	query := "SELECT * FROM building;"
	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var ans []model.Building
	for row.Next() {
		var bld model.Building
		err := row.Scan(&bld.Id, &bld.Class, &bld.NumberOfFloors, &bld.Address)
		if err != nil {
			return nil, err
		}
		ans = append(ans, bld)
	}
	return ans, nil
}

func (r *BuildingRepository) Update(bld model.Building) error {
	query := "UPDATE building SET class=$1, number_of_floors=$2, address=$3 WHERE id=$4;"
	_, err := r.db.Exec(query, bld.Class, bld.NumberOfFloors, bld.Address, bld.Id)
	return err
}

func (r *BuildingRepository) Delete(id string) error {
	query := "DELETE FROM building WHERE id=$1;"
	_, err := r.db.Exec(query, id)
	return err
}