package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

type Company struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	Description string `json:"description"`
	Country     string `json:"country"`
	City        string `json:"city"`
}

func (u *Company) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 40)),
		validation.Field(&u.Description, validation.Required),
		validation.Field(&u.Country, validation.Required),
		validation.Field(&u.City, validation.Required),
	)
}

func (u *Company) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encriptString(u.Password)
		if err != nil {
			return err
		}
		u.Password = enc
	}
	return nil
}

func (u *Company) Sanitize() {
	u.Password = ""
}

func (u *Company) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func encriptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
