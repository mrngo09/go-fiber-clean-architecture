package accountstorage

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"context"

	"gorm.io/gorm"
)

func (s *psqlStorage) FindAccount(
	ctx context.Context,
	condition map[string]interface{},
) (*accountmodel.Account, error) {
	var accountData accountmodel.Account

	if err := s.db.Where(condition).First(&accountData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, accountmodel.ErrAccountNotFound
		}
		return nil, err
	}
	return &accountData, nil
}
