package accountstorage

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"context"
)

func (s *mysqlStorage) DeleteAccount(
	ctx context.Context,
	condition map[string]interface{},
) error {
	if err := s.db.Table(accountmodel.Account{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
