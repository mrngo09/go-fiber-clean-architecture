package accountbusiness

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"context"
)

type FindAccountStorage interface {
	FindAccount(
		ctx context.Context,
		condition map[string]interface{},
	) (*accountmodel.Account, error)
}

type findBiz struct {
	store FindAccountStorage
}

func NewFindAccountBiz(store FindAccountStorage) *findBiz {
	return &findBiz{store: store}
}

func (biz *findBiz) FindAnAccount(
	ctx context.Context,
	condition map[string]interface{},
) (*accountmodel.Account, error) {
	accountData, err := biz.store.FindAccount(ctx, condition)
	if err != nil {
		return nil, err
	}

	return accountData, nil
}
