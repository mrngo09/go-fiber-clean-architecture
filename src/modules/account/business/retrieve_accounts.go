package accountbiz

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"context"
)

type RetrieveAccountsStorage interface {
	RetrieveAccount(
		c context.Context,
		condition map[string]interface{},
		pagging *accountmodel.DataPaging,
	) ([]accountmodel.Account, error)
}

type retrieveBiz struct {
	store RetrieveAccountsStorage
}

func NewRetrieveAccountsBiz(store RetrieveAccountsStorage) *retrieveBiz {
	return &retrieveBiz{store: store}
}

func (biz *retrieveBiz) RetrieveAccounts(ctx context.Context, condition map[string]interface{}, pagging *accountmodel.DataPaging) ([]accountmodel.Account, error) {
	data, err := biz.store.RetrieveAccount(ctx, condition, pagging)
	if err != nil {
		return nil, err
	}

	return data, nil
}
