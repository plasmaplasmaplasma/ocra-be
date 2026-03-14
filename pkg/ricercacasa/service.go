package ricercacasa

import (
	"errors"

	"ocra/pkg/entities"
)

type CreateRequest struct {
	TempoDiAcquisto  int32   `json:"tempo_di_acquisto"`
	Budget           float64 `json:"budget"`
	PercentualeMutuo float64 `json:"percentuale_mutuo"`
	Liquidita        float64 `json:"liquidita"`
	ClienteID        int64   `json:"cliente_id"`
	CasaID           int64   `json:"casa_id"`
}

type UpdateRequest struct {
	TempoDiAcquisto  int32   `json:"tempo_di_acquisto"`
	Budget           float64 `json:"budget"`
	PercentualeMutuo float64 `json:"percentuale_mutuo"`
	Liquidita        float64 `json:"liquidita"`
	ClienteID        int64   `json:"cliente_id"`
	CasaID           int64   `json:"casa_id"`
}

type Service interface {
	List(filter Filter) ([]entities.SearchHouse, int64, error)
	Create(req CreateRequest) (*entities.SearchHouse, error)
	Update(id int64, req UpdateRequest) (*entities.SearchHouse, error)
	Delete(id int64) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func validateRequest(clienteID, casaID int64, percentualeMutuo float64) error {
	if clienteID == 0 {
		return errors.New("cliente_id è obbligatorio")
	}
	if casaID == 0 {
		return errors.New("casa_id è obbligatorio")
	}
	if percentualeMutuo < 0 || percentualeMutuo > 100 {
		return errors.New("percentuale_mutuo deve essere tra 0 e 100")
	}
	return nil
}

func (s *service) List(filter Filter) ([]entities.SearchHouse, int64, error) {
	return s.repository.List(filter)
}

func (s *service) Create(req CreateRequest) (*entities.SearchHouse, error) {
	if err := validateRequest(req.ClienteID, req.CasaID, req.PercentualeMutuo); err != nil {
		return nil, err
	}
	rc := &entities.SearchHouse{
		TempoDiAcquisto:  req.TempoDiAcquisto,
		Budget:           req.Budget,
		PercentualeMutuo: req.PercentualeMutuo,
		Liquidita:        req.Liquidita,
		ClienteID:        req.ClienteID,
		CasaID:           req.CasaID,
	}
	if err := s.repository.Create(rc); err != nil {
		return nil, err
	}
	return s.repository.FindByID(rc.ID)
}

func (s *service) Update(id int64, req UpdateRequest) (*entities.SearchHouse, error) {
	if err := validateRequest(req.ClienteID, req.CasaID, req.PercentualeMutuo); err != nil {
		return nil, err
	}
	rc, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	rc.TempoDiAcquisto = req.TempoDiAcquisto
	rc.Budget = req.Budget
	rc.PercentualeMutuo = req.PercentualeMutuo
	rc.Liquidita = req.Liquidita
	rc.ClienteID = req.ClienteID
	rc.CasaID = req.CasaID
	if err := s.repository.Update(rc); err != nil {
		return nil, err
	}
	return s.repository.FindByID(id)
}

func (s *service) Delete(id int64) error {
	return s.repository.Delete(id)
}
