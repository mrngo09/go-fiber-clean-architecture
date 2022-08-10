package jwt

import (
	"clean-architecture-go-fiber/src/components/tokenprovider"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtProvider struct {
	secret string
}

func NewJwtProvider(secret string) *jwtProvider {
	return &jwtProvider{
		secret: secret,
	}
}

type myClaims struct {
	Payload *tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data *tokenprovider.TokenPayload, expire int) (*tokenprovider.Token, error) {
	//generate jwt
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expire)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		Token:   myToken,
		Created: time.Now(),
		Expire:  expire,
	}, nil
}

func (j *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	return claims.Payload, nil
}
