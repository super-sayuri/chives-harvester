package user

import (
	"sayuri_crypto_bot/db"
)

type CUser struct {
	db.GormModel
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Salt     string `json:"salt"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (u *CUser) HideSensitive() {
	u.Salt = "******"
	u.Password = "******"
}
