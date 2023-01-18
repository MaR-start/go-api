package sqlstore

import (
	"database/sql"

	"github.com/MaR-start/go-api/internal/app/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db                *sql.DB
	comapnyRepository *CompanyRepository
	vacancyRepository *VacancyRepository
}

// New...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Company ...
func (s *Store) Company() store.CompanyRepository {
	if s.comapnyRepository != nil {
		return s.comapnyRepository
	}

	s.comapnyRepository = &CompanyRepository{
		store: s,
	}

	return s.comapnyRepository
}

func (s *Store) Vacancy() store.VacancyRepository {
	if s.vacancyRepository != nil {
		return s.vacancyRepository
	}

	s.vacancyRepository = &VacancyRepository{
		store: s,
	}

	return s.vacancyRepository
}
