package store

import "errors"

var (
	ErrorCompanyNotFound = errors.New("Company not found")
	ErrorEmptyVacancy    = errors.New("Vacancy is empty")
)
