package cliente

import (
	"errors"

	"ocra/pkg/entities"

	"gorm.io/gorm"
)

const pageSize = 20

type Filter struct {
	Acquista           *bool
	Vende              *bool
	VendePerAcquistare *bool
	ZonaID             *int64
	SortBy             string
	SortDir            string
	Page               int
}

type Repository interface {
	List(filter Filter) ([]entities.Client, int64, error)
	FindByID(id int64) (*entities.Client, error)
	Create(cliente *entities.Client) error
	Update(cliente *entities.Client) error
	Delete(id int64) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) List(filter Filter) ([]entities.Client, int64, error) {
	var clients []entities.Client
	var total int64

	buildQuery := func() *gorm.DB {
		query := r.DB.Session(&gorm.Session{}).Model(&entities.Client{})

		if filter.Acquista != nil {
			query = query.Where("acquista = ?", *filter.Acquista)
		}
		if filter.Vende != nil {
			query = query.Where("vende = ?", *filter.Vende)
		}
		if filter.VendePerAcquistare != nil {
			query = query.Where("vende_per_acquistare = ?", *filter.VendePerAcquistare)
		}
		if filter.ZonaID != nil {
			casaSubQ := r.DB.Session(&gorm.Session{}).Model(&entities.House{}).Select("id").Where("zona_id = ?", *filter.ZonaID)
			ricercaSubQ := r.DB.Session(&gorm.Session{}).Model(&entities.SearchHouse{}).Select("cliente_id").Where("casa_id IN (?)", casaSubQ)
			query = query.Where("id IN (?)", ricercaSubQ)
		}

		return query
	}

	countQuery := buildQuery()
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sortBy := "created_at"
	validSortFields := map[string]bool{
		"nome": true, "cognome": true, "numero_di_telefono": true, "email": true,
	}
	if validSortFields[filter.SortBy] {
		sortBy = filter.SortBy
	}

	sortDir := "asc"
	if filter.SortDir == "desc" {
		sortDir = "desc"
	}

	page := filter.Page
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	dataQuery := buildQuery()
	err := dataQuery.Order(sortBy + " " + sortDir).
		Limit(pageSize).
		Offset(offset).
		Find(&clients).Error

	return clients, total, err
}

func (r *repository) FindByID(id int64) (*entities.Client, error) {
	var cliente entities.Client
	if err := r.DB.First(&cliente, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("cliente non trovato")
		}
		return nil, err
	}
	return &cliente, nil
}

func (r *repository) Create(cliente *entities.Client) error {
	return r.DB.Create(cliente).Error
}

func (r *repository) Update(cliente *entities.Client) error {
	return r.DB.Save(cliente).Error
}

func (r *repository) Delete(id int64) error {
	return r.DB.Delete(&entities.Client{}, id).Error
}
