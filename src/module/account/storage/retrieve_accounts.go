package accountstorage

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"context"
)

func (s *mysqlStorage) RetrieveAccount(
	ctx context.Context,
	condition map[string]interface{},
	pagging *accountmodel.DataPaging,
) ([]accountmodel.Account, error) {
	offset := (pagging.Page - 1) * pagging.Limit

	var data []accountmodel.Account

	err := s.db.Table(accountmodel.Account{}.TableName()).Where(condition).Count(&pagging.Total).Offset(offset).Order("id desc").Find(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}
