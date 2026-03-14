package casa

import (
	"errors"

	"ocra/pkg/entities"
)

type CreateRequest struct {
	Piano          *int16 `json:"piano"`
	NumeroDiLocali *int16 `json:"numero_di_locali"`
	NumeroDiCamere *int16 `json:"numero_di_camere"`
	NumeroDiBagni  *int16 `json:"numero_di_bagni"`
	Balcone        *bool  `json:"balcone"`
	Terrazzo       *bool  `json:"terrazzo"`
	Giardino       *bool  `json:"giardino"`
	ZonaID         int64  `json:"zona_id"`
}

type UpdateRequest struct {
	Piano          *int16 `json:"piano"`
	NumeroDiLocali *int16 `json:"numero_di_locali"`
	NumeroDiCamere *int16 `json:"numero_di_camere"`
	NumeroDiBagni  *int16 `json:"numero_di_bagni"`
	Balcone        *bool  `json:"balcone"`
	Terrazzo       *bool  `json:"terrazzo"`
	Giardino       *bool  `json:"giardino"`
	ZonaID         int64  `json:"zona_id"`
}

type Service interface {
	List(filter Filter) ([]entities.House, int64, error)
	Create(req CreateRequest) (*entities.House, error)
	Update(id int64, req UpdateRequest) (*entities.House, error)
	Delete(id int64) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) List(filter Filter) ([]entities.House, int64, error) {
	return s.repository.List(filter)
}

func (s *service) Create(req CreateRequest) (*entities.House, error) {
	if req.ZonaID == 0 {
		return nil, errors.New("zona_id è obbligatorio")
	}
	casa := &entities.House{
		Piano:          req.Piano,
		NumeroDiLocali: req.NumeroDiLocali,
		NumeroDiCamere: req.NumeroDiCamere,
		NumeroDiBagni:  req.NumeroDiBagni,
		Balcone:        req.Balcone,
		Terrazzo:       req.Terrazzo,
		Giardino:       req.Giardino,
		ZonaID:         req.ZonaID,
	}
	if err := s.repository.Create(casa); err != nil {
		return nil, err
	}
	return s.repository.FindByID(casa.ID)
}

func (s *service) Update(id int64, req UpdateRequest) (*entities.House, error) {
	if req.ZonaID == 0 {
		return nil, errors.New("zona_id è obbligatorio")
	}
	casa, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	casa.Piano = req.Piano
	casa.NumeroDiLocali = req.NumeroDiLocali
	casa.NumeroDiCamere = req.NumeroDiCamere
	casa.NumeroDiBagni = req.NumeroDiBagni
	casa.Balcone = req.Balcone
	casa.Terrazzo = req.Terrazzo
	casa.Giardino = req.Giardino
	casa.ZonaID = req.ZonaID
	if err := s.repository.Update(casa); err != nil {
		return nil, err
	}
	return s.repository.FindByID(id)
}

func (s *service) Delete(id int64) error {
	return s.repository.Delete(id)
}
