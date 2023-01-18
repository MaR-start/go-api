package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Vacancy struct {
	Id          int    `json:"id"`
	Position    string `json:"position"`
	Experience  int    `json:"experience"`
	Description string `json:"description"`
	Company_id  int    `json:"company_id"`
}

func (u *Vacancy) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Position, validation.Required),
		validation.Field(&u.Experience, validation.Required),
		validation.Field(&u.Description, validation.Required),
		validation.Field(&u.Company_id, validation.Required),
	)
}
