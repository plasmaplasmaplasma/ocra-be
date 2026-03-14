package casa

import (
	"errors"

	"ocra/pkg/entities"

	"gorm.io/gorm"
)

const pageSize = 20

type Filter struct {
	Piano          *int16
	NumeroDiLocali *int16
	NumeroDiCamere *int16
	NumeroDiBagni  *int16
	Balcone        *bool
	Terrazzo       *bool
	Giardino       *bool
	ZonaID         *int64
	SortBy         string
	SortDir        string
	Page           int
}

type Repository interface {
	List(filter Filter) ([]entities.House, int64, error)
	FindByID(id int64) (*entities.House, error)
	Create(casa *entities.House) error
	Update(casa *entities.House) error
	Delete(id int64) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) List(filter Filter) ([]entities.House, int64, error) {
	var houses []entities.House
	var total int64

	buildQuery := func() *gorm.DB {
		query := r.DB.Session(&gorm.Session{}).Model(&entities.House{})

		if filter.Piano != nil {
			query = query.Where("piano = ?", *filter.Piano)
		}
		if filter.NumeroDiLocali != nil {
			query = query.Where("numero_di_locali = ?", *filter.NumeroDiLocali)
		}
		if filter.NumeroDiCamere != nil {
			query = query.Where("numero_di_camere = ?", *filter.NumeroDiCamere)
		}
		if filter.NumeroDiBagni != nil {
			query = query.Where("numero_di_bagni = ?", *filter.NumeroDiBagni)
		}
		if filter.Balcone != nil {
			query = query.Where("balcone = ?", *filter.Balcone)
		}
		if filter.Terrazzo != nil {
			query = query.Where("terrazzo = ?", *filter.Terrazzo)
		}
		if filter.Giardino != nil {
			query = query.Where("giardino = ?", *filter.Giardino)
		}
		if filter.ZonaID != nil {
			query = query.Where("zona_id = ?", *filter.ZonaID)
		}

		return query
	}

	countQuery := buildQuery()
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sortBy := "created_at"
	validSortFields := map[string]bool{
		"piano": true, "numero_di_locali": true, "numero_di_camere": true, "numero_di_bagni": true,
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
	err := dataQuery.Preload("Zona").
		Order(sortBy + " " + sortDir).
		Limit(pageSize).
		Offset(offset).
		Find(&houses).Error

	return houses, total, err
}

func (r *repository) FindByID(id int64) (*entities.House, error) {
	var casa entities.House
	if err := r.DB.Preload("Zona").First(&casa, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("casa non trovata")
		}
		return nil, err
	}
	return &casa, nil
}

func (r *repository) Create(casa *entities.House) error {
	return r.DB.Create(casa).Error
}

func (r *repository) Update(casa *entities.House) error {
	return r.DB.Save(casa).Error
}

func (r *repository) Delete(id int64) error {
	return r.DB.Delete(&entities.House{}, id).Error
}
