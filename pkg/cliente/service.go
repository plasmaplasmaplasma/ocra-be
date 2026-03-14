package cliente

import (
	"errors"

	"ocra/pkg/entities"
)

type CreateRequest struct {
	Nome               *string `json:"nome"`
	Cognome            *string `json:"cognome"`
	NumeroDiTelefono   *string `json:"numero_di_telefono"`
	Email              *string `json:"email"`
	Acquista           bool    `json:"acquista"`
	Vende              bool    `json:"vende"`
	VendePerAcquistare bool    `json:"vende_per_acquistare"`
}

type UpdateRequest struct {
	Nome               *string `json:"nome"`
	Cognome            *string `json:"cognome"`
	NumeroDiTelefono   *string `json:"numero_di_telefono"`
	Email              *string `json:"email"`
	Acquista           bool    `json:"acquista"`
	Vende              bool    `json:"vende"`
	VendePerAcquistare bool    `json:"vende_per_acquistare"`
}

type Service interface {
	List(filter Filter) ([]entities.Client, int64, error)
	Create(req CreateRequest) (*entities.Client, error)
	Update(id int64, req UpdateRequest) (*entities.Client, error)
	Delete(id int64) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func validateMutualExclusion(vende, vendePerAcquistare bool) error {
	if vende && vendePerAcquistare {
		return errors.New("vende e vende_per_acquistare sono mutualmente esclusivi")
	}
	return nil
}

func (s *service) List(filter Filter) ([]entities.Client, int64, error) {
	return s.repository.List(filter)
}

func (s *service) Create(req CreateRequest) (*entities.Client, error) {
	if err := validateMutualExclusion(req.Vende, req.VendePerAcquistare); err != nil {
		return nil, err
	}
	cliente := &entities.Client{
		Nome:               req.Nome,
		Cognome:            req.Cognome,
		NumeroDiTelefono:   req.NumeroDiTelefono,
		Email:              req.Email,
		Acquista:           req.Acquista,
		Vende:              req.Vende,
		VendePerAcquistare: req.VendePerAcquistare,
	}
	if err := s.repository.Create(cliente); err != nil {
		return nil, err
	}
	return cliente, nil
}

func (s *service) Update(id int64, req UpdateRequest) (*entities.Client, error) {
	if err := validateMutualExclusion(req.Vende, req.VendePerAcquistare); err != nil {
		return nil, err
	}
	cliente, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	cliente.Nome = req.Nome
	cliente.Cognome = req.Cognome
	cliente.NumeroDiTelefono = req.NumeroDiTelefono
	cliente.Email = req.Email
	cliente.Acquista = req.Acquista
	cliente.Vende = req.Vende
	cliente.VendePerAcquistare = req.VendePerAcquistare
	if err := s.repository.Update(cliente); err != nil {
		return nil, err
	}
	return cliente, nil
}

func (s *service) Delete(id int64) error {
	return s.repository.Delete(id)
}
