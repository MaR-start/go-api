package store

import "github.com/MaR-start/go-api/internal/app/model"

type CompanyRepository interface {
	Create(*model.Company) error
	FindByEmail(string) (*model.Company, error)
}

type VacancyRepository interface {
	Create(*model.Vacancy) error
	FindAll(string) ([]*model.Vacancy, error)
}
