package accountmodel

import "time"

type Account struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Email     string     `json:"email" gorm:"column:email;"`
	Password  string     `json:"password" gorm:"column:password;"`
	Status    bool       `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Account) TableName() string {
	return "Account"
}
