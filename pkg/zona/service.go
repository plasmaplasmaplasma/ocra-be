package zona

import (
	"errors"
	"strings"

	"ocra/pkg/entities"
)

type CreateRequest struct {
	Nome string `json:"nome"`
}

type UpdateRequest struct {
	Nome string `json:"nome"`
}

type Service interface {
	List(page int) ([]entities.Zone, int64, error)
	Create(req CreateRequest) (*entities.Zone, error)
	Update(id int64, req UpdateRequest) (*entities.Zone, error)
	Delete(id int64) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) List(page int) ([]entities.Zone, int64, error) {
	return s.repository.List(page)
}

func (s *service) Create(req CreateRequest) (*entities.Zone, error) {
	if strings.TrimSpace(req.Nome) == "" {
		return nil, errors.New("il nome della zona è obbligatorio")
	}
	zona := &entities.Zone{Nome: req.Nome}
	if err := s.repository.Create(zona); err != nil {
		return nil, err
	}
	return zona, nil
}

func (s *service) Update(id int64, req UpdateRequest) (*entities.Zone, error) {
	if strings.TrimSpace(req.Nome) == "" {
		return nil, errors.New("il nome della zona è obbligatorio")
	}
	zona, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	zona.Nome = req.Nome
	if err := s.repository.Update(zona); err != nil {
		return nil, err
	}
	return zona, nil
}

func (s *service) Delete(id int64) error {
	return s.repository.Delete(id)
}
