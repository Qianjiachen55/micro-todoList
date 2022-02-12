package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PasswordDigest string

}

const (
	PassWordCost = 12
)


//加密
func (user *User) SetPassword (password string) error {
	bytes,err :=bcrypt.GenerateFromPassword([]byte(password), PassWordCost)

	if err != nil{
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

//解密

func (user *User)CheckPassword (password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.PasswordDigest))
	return  err == nil
}
