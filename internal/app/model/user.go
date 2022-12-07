package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

//Validation
func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(
			&u.Email,            //проверка значения
			validation.Required, //обязательное для заполнения
			is.Email),           //это поле является email

		validation.Field(
			&u.Password, //проверка значения
			validation.By(requiredIf(u.EncryptedPassword == "")), //обязательное для заполнения
			validation.Length(6, 100)),                           //это поле является email

	)
}

//BeforeCreate
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
