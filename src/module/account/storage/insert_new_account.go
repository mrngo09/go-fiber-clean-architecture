package accountstorage

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"context"
)

func (s *psqlStorage) CreateAccount(ctx context.Context, data *accountmodel.Account) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
