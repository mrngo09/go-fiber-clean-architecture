package accountstorage

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"context"
)

func (s *psqlStorage) UpdateAccount(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *accountmodel.Account,
) error {
	if err := s.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
