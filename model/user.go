package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64  `json:"id" xorm:"pk autoincr"`
	Username string `json:"username"`
	Password string `json:"-"`
	Created  string `json:"created" xorm:"created"`
	Updated  string `json:"updated" xorm:"updated"`
	Deleted  string `json:"deleted" xorm:"deleted"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}
