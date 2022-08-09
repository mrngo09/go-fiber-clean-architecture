package accountbiz

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"context"
)

type DeleteAccountStorage interface {
	DeleteAccount(
		ctx context.Context,
		condition map[string]interface{},
	) error

	FindAccount(
		ctx context.Context,
		condition map[string]interface{},
	) (*accountmodel.Account, error)
}

type deleteBiz struct {
	store DeleteAccountStorage
}

func NewDeleteAccountStorage(store DeleteAccountStorage) *deleteBiz {
	return &deleteBiz{store: store}
}

func (biz *deleteBiz) DeleteAccount(
	ctx context.Context,
	condition map[string]interface{},
) error {
	//Find account by conditions
	_, err := biz.store.FindAccount(ctx, condition)

	if err != nil {
		return err
	}

	//Call storage to delete account
	if err := biz.store.DeleteAccount(ctx, condition); err != nil {
		return err
	}
	return nil
}
