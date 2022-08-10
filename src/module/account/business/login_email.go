package accountbiz

import (
	"clean-architecture-go-fiber/src/components/tokenprovider"
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	"context"
)

type loginBiz struct {
	storeUser     FindAccountStorage
	tokenProvider tokenprovider.TokenProvider
	expire        int
}

func NewLoginBiz(
	storeUser FindAccountStorage,
	tokenProvider tokenprovider.TokenProvider,
	expire int,
) *loginBiz {
	return &loginBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		expire:        expire,
	}
}

func (biz *loginBiz) UserLogin(ctx context.Context, data *accountmodel.UserLogin) (*accountmodel.CurrentProfile, error) {

	userDB, err := biz.storeUser.FindAccount(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, err
	}

	// err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(data.Password))

	payload := tokenprovider.TokenPayload{
		UserId: userDB.Id,
		Role:   "admin",
	}
	token, err := biz.tokenProvider.Generate(&payload, biz.expire)

	if err != nil {
		return nil, err
	}

	account := accountmodel.UserLogined(token.Token, nil, userDB)
	return account, nil
}
