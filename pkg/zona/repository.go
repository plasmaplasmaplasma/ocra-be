package zona

import (
	"errors"

	"ocra/pkg/entities"

	"gorm.io/gorm"
)

const pageSize = 20

type Repository interface {
	List(page int) ([]entities.Zone, int64, error)
	FindByID(id int64) (*entities.Zone, error)
	Create(zona *entities.Zone) error
	Update(zona *entities.Zone) error
	Delete(id int64) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) List(page int) ([]entities.Zone, int64, error) {
	var zones []entities.Zone
	var total int64

	countQuery := r.DB.Session(&gorm.Session{}).Model(&entities.Zone{})
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	dataQuery := r.DB.Session(&gorm.Session{}).Model(&entities.Zone{})
	err := dataQuery.
		Order("zones.created_at asc").
		Limit(pageSize).
		Offset(offset).
		Find(&zones).Error
	return zones, total, err
}

func (r *repository) FindByID(id int64) (*entities.Zone, error) {
	var zona entities.Zone
	if err := r.DB.First(&zona, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("zona non trovata")
		}
		return nil, err
	}
	return &zona, nil
}

func (r *repository) Create(zona *entities.Zone) error {
	return r.DB.Create(zona).Error
}

func (r *repository) Update(zona *entities.Zone) error {
	return r.DB.Save(zona).Error
}

func (r *repository) Delete(id int64) error {
	return r.DB.Delete(&entities.Zone{}, id).Error
}
