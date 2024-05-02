package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/TandDA/coursedb/internal/handler"
	"github.com/TandDA/coursedb/internal/repository"
	"github.com/TandDA/coursedb/internal/service"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:123@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	doMigration(db)
	repo := repository.NewRepository(db)
	srvc := service.NewService(repo)
	hndl := handler.NewHandler(srvc, db)

	hndl.Start()
}

func doMigration(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		fmt.Println(err)
		return
	}
	m.Up()
}
