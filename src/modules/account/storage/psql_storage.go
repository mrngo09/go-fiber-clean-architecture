package accountstorage

import "gorm.io/gorm"

type psqlStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(db *gorm.DB) *psqlStorage {
	return &psqlStorage{db: db}
}
