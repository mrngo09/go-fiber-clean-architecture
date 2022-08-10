package accountstorage

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"context"

	"gorm.io/gorm"
)

type psqlStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(db *gorm.DB) *psqlStorage {
	return &psqlStorage{db: db}
}

type UserStore interface {
	FindUser(ctx context.Context,
		condition map[string]interface{},
	) (*accountmodel.Account, error)
}
