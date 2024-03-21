package repository

import (
	"database/sql"

	"github.com/TandDA/coursedb/internal/model"
)

type Repository struct {
	Building
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Building: NewBuildingRepository(db),
	}
}

type Building interface {
	Save(b model.Building) (string, error)
}