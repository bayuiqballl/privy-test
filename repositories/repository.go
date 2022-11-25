package repositories

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

type RepositoriesInterface interface {
	CakeRepository
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{db: db}
}
