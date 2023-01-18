package store

type Store interface {
	Company() CompanyRepository
	Vacancy() VacancyRepository
}
