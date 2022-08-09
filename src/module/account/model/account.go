package accountmodel

import (
	"errors"
	"time"
)

var (
	ErrPasswordCannotBeBlank       = errors.New("Account can not be blank")
	ErrAccountNotFound             = errors.New("Account not found")
	ErrCannotUpdateFinishedAccount = errors.New("Can not update finished account")
)

type Account struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Email     string     `json:"email" gorm:"column:email;"`
	Password  string     `json:"password" gorm:"column:password;"`
	Status    bool       `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Account) TableName() string {
	return "accounts"
}

type DataPaging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *DataPaging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}
}
