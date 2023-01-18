package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/MaR-start/go-api/internal/app/model"
	"github.com/MaR-start/go-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	sessionName = "practice"
)

var (
	errorIncorrectEmailOrPassword = errors.New("Incorrect password or email")
)

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configRouter() {
	s.router.HandleFunc("/company/register", s.CompanyCreate()).Methods("POST")
	s.router.HandleFunc("/company/login", s.CompanyLogin()).Methods("POST")
	s.router.HandleFunc("/vacancy/create", s.VacancyCreate()).Methods("POST")
	s.router.HandleFunc("/vacancy/find", s.VacancyFind()).Methods("POST")
}

func (s *server) CompanyCreate() http.HandlerFunc {

	type request struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		Description string `json:"description"`
		Country     string `json:"country"`
		City        string `json:"city"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.Company{
			Name:        req.Name,
			Email:       req.Email,
			Password:    req.Password,
			Description: req.Description,
			Country:     req.Country,
			City:        req.City,
		}

		if err := s.store.Company().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) CompanyLogin() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.Company().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["company_id"] = u.Id

		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *server) VacancyCreate() http.HandlerFunc {

	type request struct {
		Position    string `json:"position"`
		Experience  string `json:"experience"`
		Description string `json:"description"`
		Company_id  string `json:"company_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		a := req.Experience
		ExperienceReq, _ := strconv.Atoi(a)

		b := req.Company_id
		Company_idReq, _ := strconv.Atoi(b)

		u := &model.Vacancy{
			Position:    req.Position,
			Experience:  ExperienceReq,
			Description: req.Description,
			Company_id:  Company_idReq,
		}

		if err := s.store.Vacancy().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) VacancyFind() http.HandlerFunc {

	type request struct {
		Company_id string `json:"company_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.Vacancy().FindAll(req.Company_id)
		if err != nil {
			s.error(w, r, http.StatusBadGateway, err)
			return
		}

		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
