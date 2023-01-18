package sqlstore

import (
	"database/sql"

	"github.com/MaR-start/go-api/internal/app/model"
	"github.com/MaR-start/go-api/internal/app/store"
)

type VacancyRepository struct {
	store *Store
}

func (r *VacancyRepository) Create(u *model.Vacancy) error {

	if err := u.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO vacancies (position, experience, description, company_id) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Position,
		u.Experience,
		u.Description,
		u.Company_id,
	).Scan(&u.Id)
}

func (r *VacancyRepository) FindAll(Company_id string) ([]*model.Vacancy, error) {

	var vacancyList []*model.Vacancy

	rows, err := r.store.db.Query("SELECT * FROM vacancies WHERE company_id = $1", Company_id)

	for rows.Next() {
		f := &model.Vacancy{}
		err = rows.Scan(&f.Id, &f.Position, &f.Experience, &f.Description, &f.Company_id)
		if err == sql.ErrNoRows {
			return nil, store.ErrorEmptyVacancy
		}

		vacancyList = append(vacancyList, f)
	}

	return vacancyList, nil
}
