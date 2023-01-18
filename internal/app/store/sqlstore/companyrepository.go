package sqlstore

import (
	"database/sql"

	"github.com/MaR-start/go-api/internal/app/model"
	"github.com/MaR-start/go-api/internal/app/store"
)

type CompanyRepository struct {
	store *Store
}

func (r *CompanyRepository) Create(u *model.Company) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO companies (name, email, password, description, country, city) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		u.Name,
		u.Email,
		u.Password,
		u.Description,
		u.Country,
		u.City,
	).Scan(&u.Id)
}

func (r *CompanyRepository) FindByEmail(email string) (*model.Company, error) {
	u := &model.Company{}

	if err := r.store.db.QueryRow(
		"SELECT id, name, email, password, description, country, city FROM companies WHERE email = $1", email).Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Description,
		&u.Country,
		&u.City,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrorCompanyNotFound
		}
	}

	return u, nil
}
