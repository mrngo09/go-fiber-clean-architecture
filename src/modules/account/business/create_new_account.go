package accountbiz

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"context"
	"errors"
)

type CreateAccountStorage interface {
	CreateAccount(ctx context.Context, data *accountmodel.Account) error
}

type createBusiness struct {
	store CreateAccountStorage
}

func NewCreateAccountbiz(store CreateAccountStorage) *createBusiness {
	return &createBusiness{store: store}
}

func (business *createBusiness) CreateNewAccount(ctx context.Context, data *accountmodel.Account) error {
	if data.Email == "" {
		return errors.New("Email can not be blank.")
	}

	if data.Password == "" {
		return errors.New("Email can not be blank.")
	}

	data.Status = true
	err := business.store.CreateAccount(ctx, data)

	if err != nil {
		return err
	}

	return err
}
