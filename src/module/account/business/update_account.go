package accountbiz

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"context"
)

type UpdateAccountStorage interface {
	UpdateAccount(
		c context.Context,
		condition map[string]interface{},
		dataUpdate *accountmodel.Account,
	) error
}

type updateBiz struct {
	store UpdateAccountStorage
}

func NewUpdateAccountBiz(store UpdateAccountStorage) *updateBiz {
	return &updateBiz{store: store}
}

func (biz *updateBiz) UpdateAccount(
	c context.Context,
	condition map[string]interface{},
	dataUpdate *accountmodel.Account,
) error {
	err := biz.UpdateAccount(c, condition, dataUpdate)

	if err != nil {
		return err
	}

	return nil
}
