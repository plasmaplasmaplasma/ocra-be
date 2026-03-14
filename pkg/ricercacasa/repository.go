package ricercacasa

import (
	"errors"

	"ocra/pkg/entities"

	"gorm.io/gorm"
)

const pageSize = 20

type Filter struct {
	TempoDiAcquistoFrom  *int32
	TempoDiAcquistoTo    *int32
	BudgetFrom           *float64
	BudgetTo             *float64
	PercentualeMutuoFrom *float64
	PercentualeMutuoTo   *float64
	LiquiditaFrom        *float64
	LiquiditaTo          *float64
	ClienteID            *int64
	Piano                *int16
	NumeroDiLocali       *int16
	NumeroDiCamere       *int16
	NumeroDiBagni        *int16
	Balcone              *bool
	Terrazzo             *bool
	Giardino             *bool
	ZonaID               *int64
	Page                 int
}

type Repository interface {
	List(filter Filter) ([]entities.SearchHouse, int64, error)
	FindByID(id int64) (*entities.SearchHouse, error)
	Create(rc *entities.SearchHouse) error
	Update(rc *entities.SearchHouse) error
	Delete(id int64) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func applyFilters(query *gorm.DB, filter Filter) *gorm.DB {
	if filter.TempoDiAcquistoFrom != nil {
		query = query.Where("tempo_di_acquisto >= ?", *filter.TempoDiAcquistoFrom)
	}
	if filter.TempoDiAcquistoTo != nil {
		query = query.Where("tempo_di_acquisto <= ?", *filter.TempoDiAcquistoTo)
	}
	if filter.BudgetFrom != nil {
		query = query.Where("budget >= ?", *filter.BudgetFrom)
	}
	if filter.BudgetTo != nil {
		query = query.Where("budget <= ?", *filter.BudgetTo)
	}
	if filter.PercentualeMutuoFrom != nil {
		query = query.Where("percentuale_mutuo >= ?", *filter.PercentualeMutuoFrom)
	}
	if filter.PercentualeMutuoTo != nil {
		query = query.Where("percentuale_mutuo <= ?", *filter.PercentualeMutuoTo)
	}
	if filter.LiquiditaFrom != nil {
		query = query.Where("liquidita >= ?", *filter.LiquiditaFrom)
	}
	if filter.LiquiditaTo != nil {
		query = query.Where("liquidita <= ?", *filter.LiquiditaTo)
	}
	if filter.ClienteID != nil {
		query = query.Where("cliente_id = ?", *filter.ClienteID)
	}

	hasCasaFilter := filter.Piano != nil || filter.NumeroDiLocali != nil ||
		filter.NumeroDiCamere != nil || filter.NumeroDiBagni != nil ||
		filter.Balcone != nil || filter.Terrazzo != nil ||
		filter.Giardino != nil || filter.ZonaID != nil

	if hasCasaFilter {
		casaQuery := query.Session(&gorm.Session{}).Model(&entities.House{}).Select("id")
		if filter.Piano != nil {
			casaQuery = casaQuery.Where("piano = ?", *filter.Piano)
		}
		if filter.NumeroDiLocali != nil {
			casaQuery = casaQuery.Where("numero_di_locali = ?", *filter.NumeroDiLocali)
		}
		if filter.NumeroDiCamere != nil {
			casaQuery = casaQuery.Where("numero_di_camere = ?", *filter.NumeroDiCamere)
		}
		if filter.NumeroDiBagni != nil {
			casaQuery = casaQuery.Where("numero_di_bagni = ?", *filter.NumeroDiBagni)
		}
		if filter.Balcone != nil {
			casaQuery = casaQuery.Where("balcone = ?", *filter.Balcone)
		}
		if filter.Terrazzo != nil {
			casaQuery = casaQuery.Where("terrazzo = ?", *filter.Terrazzo)
		}
		if filter.Giardino != nil {
			casaQuery = casaQuery.Where("giardino = ?", *filter.Giardino)
		}
		if filter.ZonaID != nil {
			casaQuery = casaQuery.Where("zona_id = ?", *filter.ZonaID)
		}
		query = query.Where("casa_id IN (?)", casaQuery)
	}

	return query
}

func (r *repository) List(filter Filter) ([]entities.SearchHouse, int64, error) {
	var results []entities.SearchHouse
	var total int64

	baseQuery := applyFilters(r.DB.Model(&entities.SearchHouse{}), filter)

	countQuery := baseQuery.Session(&gorm.Session{}).Order("").Limit(-1).Offset(-1)
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	page := filter.Page
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	err := baseQuery.
		Preload("Cliente").
		Preload("Casa").
		Preload("Casa.Zona").
		Order("search_houses.created_at asc").
		Limit(pageSize).
		Offset(offset).
		Find(&results).Error

	return results, total, err
}

func (r *repository) FindByID(id int64) (*entities.SearchHouse, error) {
	var rc entities.SearchHouse
	if err := r.DB.Preload("Cliente").Preload("Casa").Preload("Casa.Zona").First(&rc, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("ricerca casa non trovata")
		}
		return nil, err
	}
	return &rc, nil
}

func (r *repository) Create(rc *entities.SearchHouse) error {
	return r.DB.Create(rc).Error
}

func (r *repository) Update(rc *entities.SearchHouse) error {
	return r.DB.Save(rc).Error
}

func (r *repository) Delete(id int64) error {
	return r.DB.Delete(&entities.SearchHouse{}, id).Error
}
