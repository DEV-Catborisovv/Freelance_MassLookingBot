package storage

import "github.com/jmoiron/sqlx"

type IStorage interface{}

type Storage struct {
	IStorage
	db *sqlx.DB
}

func NewStorage() *Storage {
	return &Storage{}
}
